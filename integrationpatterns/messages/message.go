package messages

import (
	"context"
	"fmt"
)

// Message data that is to be transmitted via a messaging system.
type Message struct {
	Code, Value string
}

type Sender struct {
	messageStream chan<- Message
}
type Receiver struct {
	started       bool
	messageStream <-chan Message
	auditStream   chan Message
}

func NewMessageSender(messageChannel chan<- Message) Sender {
	return Sender{
		messageStream: messageChannel,
	}
}

func NewMessageReceiver(messageChannel <-chan Message) Receiver {
	newReceiver := Receiver{
		messageStream: messageChannel,
		auditStream:   make(chan Message),
	}

	return newReceiver
}

func (s Sender) Send(ctx context.Context, message Message) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("context finalized: %w", ctx.Err())
	default:
		go func() {
			s.messageStream <- message
		}()
	}

	return nil
}

func (r Receiver) Start(ctx context.Context) {
	if r.started {
		return
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case newMessage, ok := <-r.messageStream:
				if !ok {
					return
				}

				go r.doAudit(ctx, newMessage)
			}
		}
	}()

	r.started = true
}

func (r Receiver) doAudit(ctx context.Context, message Message) {
	select {
	case <-ctx.Done():
		return
	default:
		r.auditStream <- message
	}
}

func (r Receiver) Audit() <-chan Message {
	return r.auditStream
}
