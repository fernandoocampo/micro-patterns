package messages

import (
	"context"
	"log"
)

type Reply struct {
	Err     error
	Message string
}

type Request struct {
	ID   int
	Code string
}

// RequestReply defines a request reply message.
type RequestReply struct {
	Data         Request
	ReplyAddress chan Reply
}

type Replier struct {
	requestStream chan RequestReply
}

type Requester struct {
	requestStream chan<- RequestReply
	replyStream   chan Reply
}

const acceptedMessage = "accepted"

func NewRequestReplyReceiver() *Replier {
	newReplier := Replier{
		requestStream: make(chan RequestReply),
	}

	return &newReplier
}

func NewRequestReplySender(stream chan<- RequestReply) *Requester {
	requester := new(Requester)
	requester.requestStream = stream
	requester.replyStream = make(chan Reply)

	return requester
}

func (r *Requester) Send(ctx context.Context, request Request) {
	newRequest := RequestReply{
		Data:         request,
		ReplyAddress: r.replyStream,
	}

	go func(requestReply RequestReply) {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Println("error", ctx.Err().Error())
			}
			return
		case r.requestStream <- requestReply:
		}
	}(newRequest)
}

func (r *Requester) ReplyStream() <-chan Reply {
	return r.replyStream
}

func (r *Replier) RequestStream() chan<- RequestReply {
	return r.requestStream
}

func (r *Replier) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request, ok := <-r.requestStream:
			if !ok {
				log.Println("reply stream was closed unexpectedly")
				return
			}
			r.processRequest(request)
		}
	}
}

func (r *Replier) processRequest(request RequestReply) {
	log.Println("processing request", request)
	go func(request RequestReply) {
		log.Println("do something with the request")
		request.ReplyAddress <- Reply{
			Err:     nil,
			Message: acceptedMessage,
		}
	}(request)
}
