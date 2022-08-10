package messages_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages"
	"github.com/stretchr/testify/assert"
)

func TestCommandMessageCreate(t *testing.T) {
	expectedOrder := messages.Order{
		ID:       1234,
		Amount:   23.45,
		Location: "maria de los angeles",
	}
	cmdMessage := messages.CommandMessage{
		ID:   messages.Create,
		Name: "create_order",
		Parameters: []messages.Parameter{
			{
				Name:  "ID",
				Value: 1234,
			},
			{
				Name:  "amount",
				Value: 23.45,
			},
			{
				Name:  "location",
				Value: "maria de los angeles",
			},
		},
	}

	ctx := context.TODO()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	messageChannel := make(chan messages.CommandMessage)
	sender := messages.NewCommandMessageSender(messageChannel)
	receiver := messages.NewCommandMessageReceiver(messageChannel)

	receiver.Start(ctx)

	// When
	err := sender.Send(ctx, cmdMessage)

	var closed bool

	var newMessage messages.Order

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			t.Log(ctx.Err())
		}
	case newMessage, closed = <-receiver.Audit():
		if !closed {
			t.Log("receive audit channel was closed")
		}
	}
	// Then
	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, newMessage)

}
