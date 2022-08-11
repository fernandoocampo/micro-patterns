package messages_test

import (
	"context"
	"testing"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages"
	"github.com/stretchr/testify/assert"
)

func TestEventMessage(t *testing.T) {
	// Given
	event := messages.PriceEvent{
		OldPrice:  3.45,
		NewPrice:  4.23,
		ProductID: "AS2345Z1",
	}

	ctx := context.TODO()

	subscribers := []*messages.PriceSubscriber{
		messages.NewPriceSubscriber(),
		messages.NewPriceSubscriber(),
		messages.NewPriceSubscriber(),
	}

	topic := messages.NewPricesTopic()

	for _, s := range subscribers {
		s.Subscribe(ctx, topic)
	}

	publisher := messages.NewPublisher(topic)
	publisher.Publish(ctx, event)

	for _, s := range subscribers {
		assert.Equal(t, event, <-s.Audit())
	}
}
