package messages

import "context"

type Subscriber interface {
	Notify(context.Context, PriceEvent)
}

// PriceEvent event contains data to signal or announce something. e.g. Price change event.
type PriceEvent struct {
	OldPrice, NewPrice float32
	ProductID          string
}

type PriceTopic struct {
	eventStream chan PriceEvent
	subscribers []Subscriber
}

type PriceSubscriber struct {
	auditStream chan PriceEvent
}

type PricePublisher struct {
	topic *PriceTopic
}

func NewPriceSubscriber() *PriceSubscriber {
	priceSubscriber := PriceSubscriber{
		auditStream: make(chan PriceEvent),
	}

	return &priceSubscriber
}

func NewPricesTopic() *PriceTopic {
	newPriceTopic := PriceTopic{
		eventStream: make(chan PriceEvent),
	}

	return &newPriceTopic
}

func NewPublisher(topic *PriceTopic) *PricePublisher {
	newPublisher := PricePublisher{
		topic: topic,
	}

	return &newPublisher
}

func (p *PriceSubscriber) Subscribe(ctx context.Context, topic *PriceTopic) {
	topic.addSubscriber(p)
}

func (p *PriceSubscriber) Notify(ctx context.Context, event PriceEvent) {
	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			p.auditStream <- event
		}
	}()
}

func (p *PriceSubscriber) Audit() <-chan PriceEvent {
	return p.auditStream
}

func (p *PricePublisher) Publish(ctx context.Context, event PriceEvent) {
	p.topic.publish(ctx, event)
}

func (p *PriceTopic) publish(ctx context.Context, event PriceEvent) {
	for idx := range p.subscribers {
		subscriber := p.subscribers[idx]
		go func() {
			select {
			case <-ctx.Done():
				return
			default:
				subscriber.Notify(ctx, event)
			}
		}()
	}
}

func (p *PriceTopic) addSubscriber(newSubscriber Subscriber) {
	p.subscribers = append(p.subscribers, newSubscriber)
}
