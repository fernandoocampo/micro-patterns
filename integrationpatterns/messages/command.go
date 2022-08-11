package messages

import (
	"context"
	"fmt"
	"log"
)

type CommandID int32

// Parameter define command message parameters.
type Parameter struct {
	Name  string
	Value interface{}
}

// CommandMessage contains the name of the command and the parameters that must be a value with json format.
type CommandMessage struct {
	ID         CommandID
	Name       string
	Parameters []Parameter
}

// Order contains order data.
type Order struct {
	ID       int
	Amount   float32
	Location string
}

type CommandSender struct {
	messageStream chan<- CommandMessage
}

type CommandReceiver struct {
	started       bool
	messageStream <-chan CommandMessage
	auditStream   chan Order
}

// command ids.
const (
	Create CommandID = iota
	Update
	Delete
)

// order field names.
const (
	idField       = "ID"
	amountField   = "amount"
	locationField = "location"
)

func NewCommandMessageSender(messageChannel chan<- CommandMessage) CommandSender {
	return CommandSender{
		messageStream: messageChannel,
	}
}

func NewCommandMessageReceiver(messageChannel <-chan CommandMessage) CommandReceiver {
	newReceiver := CommandReceiver{
		messageStream: messageChannel,
		auditStream:   make(chan Order),
	}

	return newReceiver
}

func (s CommandSender) Send(ctx context.Context, message CommandMessage) error {
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

func (r CommandReceiver) Start(ctx context.Context) {
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

				r.processCommand(ctx, newMessage)
			}
		}
	}()

	r.started = true
}

func (r CommandReceiver) processCommand(ctx context.Context, message CommandMessage) {
	switch message.ID {
	case Create:
		log.Println("creating order")

		newOrder := r.transformCommandToOrder(message.Parameters)
		r.createOrder(ctx, newOrder)
	default:
		log.Println("we not support this kind of command", message)
	}
}

func (r CommandReceiver) transformCommandToOrder(parameters []Parameter) Order {
	var result Order

	for idx := range parameters {
		switch parameters[idx].Name {
		case idField:
			result.setID(parameters[idx].Value)
		case amountField:
			result.setAmount(parameters[idx].Value)
		case locationField:
			result.setLocation(parameters[idx].Value)
		default:
			log.Println("unknown field", parameters[idx])
		}
	}

	return result
}

func (r CommandReceiver) createOrder(ctx context.Context, order Order) {
	log.Println("creating order", order)

	go r.doAudit(ctx, order)
}

func (r CommandReceiver) doAudit(ctx context.Context, order Order) {
	select {
	case <-ctx.Done():
		return
	default:
		r.auditStream <- order
	}
}

func (r CommandReceiver) Audit() <-chan Order {
	return r.auditStream
}

func (o *Order) setID(value interface{}) {
	v, ok := value.(int)
	if !ok {
		return
	}

	o.ID = v
}

func (o *Order) setAmount(value interface{}) {
	floatValue, ok := value.(float64)
	if !ok {
		log.Printf("unexpected value: %v with type: %T", value, value)

		return
	}

	o.Amount = float32(floatValue)
}

func (o *Order) setLocation(value interface{}) {
	v, ok := value.(string)
	if !ok {
		return
	}

	o.Location = v
}
