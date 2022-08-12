package correlations

import (
	"context"
	"log"
)

type Replier struct {
	requestChannel chan Request
	replyChannel   chan Reply
}

type Reply struct {
	CorrelationID, Message string
	Err                    error
}

const acceptedMessage = "accepted"

func NewRequestReplyReceiver(requestChannel chan Request, replyChannel chan Reply) *Replier {
	newReplier := Replier{
		replyChannel:   replyChannel,
		requestChannel: requestChannel,
	}

	return &newReplier
}

func (r *Replier) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request, ok := <-r.requestChannel:
			if !ok {
				log.Println("reply stream was closed unexpectedly")

				return
			}

			r.processRequest(request)
		}
	}
}

func (r *Replier) processRequest(request Request) {
	log.Println("processing request", request)

	go func(request Request) {
		log.Println("do something with the request")
		r.replyChannel <- Reply{
			CorrelationID: request.CorrelationID,
			Err:           nil,
			Message:       acceptedMessage,
		}
	}(request)
}
