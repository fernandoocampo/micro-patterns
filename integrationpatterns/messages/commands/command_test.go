package commands_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages/commands"
	"github.com/stretchr/testify/assert"
)

func TestCommandMessageCreate(t *testing.T) {
	t.Parallel()

	expectedOrder := commands.Order{
		ID:       1234,
		Amount:   23.45,
		Location: "maria de los angeles",
	}
	cmdMessage := commands.CommandMessage{
		ID:   commands.Create,
		Name: "create_order",
		Parameters: []commands.Parameter{
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

	messageChannel := make(chan commands.CommandMessage)
	sender := commands.NewCommandMessageSender(messageChannel)
	receiver := commands.NewCommandMessageReceiver(messageChannel)

	receiver.Start(ctx)

	// When
	err := sender.Send(ctx, cmdMessage)

	var closed bool

	var newMessage commands.Order

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
