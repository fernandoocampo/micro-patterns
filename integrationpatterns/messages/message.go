package messages

import "context"

// Message data that is to be transmitted via a messaging system.
type Message struct {
	Code, Value string
}

type sender struct {
	messageStream chan<- Message
}
type receiver struct {
	started       bool
	messageStream <-chan Message
	auditStream   chan Message
}

func NewMessageSender(messageChannel chan<- Message) sender {
	return sender{
		messageStream: messageChannel,
	}
}

func NewMessageReceiver(messageChannel <-chan Message) receiver {
	newReceiver := receiver{
		messageStream: messageChannel,
		auditStream:   make(chan Message),
	}
	return newReceiver
}

func (s sender) Send(ctx context.Context, message Message) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		go func() {
			s.messageStream <- message
		}()
	}

	return nil
}

func (r receiver) Start(ctx context.Context) {
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
}

func (r receiver) doAudit(ctx context.Context, message Message) {
	select {
	case <-ctx.Done():
		return
	default:
		r.auditStream <- message
	}
}

func (r receiver) Audit() <-chan Message {
	return r.auditStream
}
