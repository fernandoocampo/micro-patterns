package correlations

import (
	"context"
	"log"
	"sync"
)

type RequestStatus struct {
	Message string
}

type Requester struct {
	requestChannel chan Request
	replyChannel   chan Reply
	auditStream    chan RequestStatus
	states         map[string]RequestStatus
	lock           sync.Mutex
}

type Request struct {
	CorrelationID, Code string
}

const (
	availableBuffer = 2
)

const (
	acceptedState  = "accepted"
	processedState = "processed"
)

func NewRequestReplySender(requestChannel chan Request, replyChannel chan Reply) *Requester {
	requester := Requester{
		replyChannel:   replyChannel,
		requestChannel: requestChannel,
		auditStream:    make(chan RequestStatus, availableBuffer),
		states:         make(map[string]RequestStatus),
	}

	return &requester
}

func (r *Requester) Send(ctx context.Context, request Request) {
	go func(request Request) {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Println("error", ctx.Err().Error())
			}

			return
		case r.requestChannel <- request:
			update := RequestStatus{
				Message: acceptedState,
			}
			r.addStatus(request.CorrelationID, update)
			r.doAudit(ctx, update)
		}
	}(request)
}

func (r *Requester) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case reply, ok := <-r.replyChannel:
			if !ok {
				log.Println("reply stream was closed unexpectedly")

				return
			}

			r.processReply(ctx, reply)
		}
	}
}

func (r *Requester) processReply(ctx context.Context, reply Reply) {
	log.Println("processing reply", reply)
	log.Println("do something with the response")

	update := r.updateState(reply.CorrelationID, processedState)
	r.doAudit(ctx, update)
}

func (r *Requester) doAudit(ctx context.Context, update RequestStatus) {
	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			r.auditStream <- update
		}
	}()
}

func (r *Requester) addStatus(key string, value RequestStatus) {
	r.lock.Lock()
	{
		r.states[key] = value
	}
	r.lock.Unlock()
}

func (r *Requester) updateState(key string, newStatus string) RequestStatus {
	r.lock.Lock()
	defer r.lock.Unlock()

	update, ok := r.states[key]
	if !ok {
		log.Println("request was not requested", key)

		return RequestStatus{}
	}

	update.Message = newStatus
	r.states[key] = update

	return update
}

func (r *Requester) Audit() <-chan RequestStatus {
	return r.auditStream
}
