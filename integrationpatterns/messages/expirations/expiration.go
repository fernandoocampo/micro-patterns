package expirations

import (
	"context"
	"fmt"
	"time"
)

// Message data that is to be transmitted via a messaging system.
type Message struct {
	Deadline int64
	Value    string
}

type Resolution struct {
	Processed bool
}

type Sender struct {
	messageStream chan<- Message
}

type Receiver struct {
	started       bool
	messageStream <-chan Message
	auditStream   chan Resolution
}

func NewMessageSender(messageChannel chan<- Message) Sender {
	return Sender{
		messageStream: messageChannel,
	}
}

func NewMessageReceiver(messageChannel <-chan Message) Receiver {
	newReceiver := Receiver{
		messageStream: messageChannel,
		auditStream:   make(chan Resolution),
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

				resolution := processMessage((newMessage))

				go r.doAudit(ctx, resolution)
			}
		}
	}()

	r.started = true
}

func processMessage(message Message) Resolution {
	var result Resolution

	if now := time.Now().Unix(); now <= message.Deadline {
		result.Processed = true
	}

	return result
}

func (r Receiver) doAudit(ctx context.Context, message Resolution) {
	select {
	case <-ctx.Done():
		return
	default:
		r.auditStream <- message
	}
}

func (r Receiver) Audit() <-chan Resolution {
	return r.auditStream
}
