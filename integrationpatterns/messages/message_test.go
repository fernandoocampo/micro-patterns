package messages_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	t.Parallel()
	// Given
	message := messages.Message{
		Code:  "200",
		Value: "Success",
	}
	expectedMessage := messages.Message{
		Code:  "200",
		Value: "Success",
	}
	ctx := context.TODO()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	messageChannel := make(chan messages.Message)
	sender := messages.NewMessageSender(messageChannel)
	receiver := messages.NewMessageReceiver(messageChannel)

	receiver.Start(ctx)

	// When
	err := sender.Send(ctx, message)

	var closed bool

	var newMessage messages.Message

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			t.Log(ctx.Err())
		}
	case newMessage, closed = <-receiver.Audit():
		if !closed {
			t.Log("receivir audit channel was closed")
		}
	}
	// Then
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, newMessage)
}
