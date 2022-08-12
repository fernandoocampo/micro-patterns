package events_test

import (
	"context"
	"testing"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages/events"
	"github.com/stretchr/testify/assert"
)

func TestEventMessage(t *testing.T) {
	t.Parallel()
	// Given
	event := events.PriceEvent{
		OldPrice:  3.45,
		NewPrice:  4.23,
		ProductID: "AS2345Z1",
	}

	ctx := context.TODO()

	subscribers := []*events.PriceSubscriber{
		events.NewPriceSubscriber(),
		events.NewPriceSubscriber(),
		events.NewPriceSubscriber(),
	}

	topic := events.NewPricesTopic()

	for _, s := range subscribers {
		s.Subscribe(ctx, topic)
	}

	publisher := events.NewPublisher(topic)
	publisher.Publish(ctx, event)

	for _, s := range subscribers {
		assert.Equal(t, event, <-s.Audit())
	}
}
