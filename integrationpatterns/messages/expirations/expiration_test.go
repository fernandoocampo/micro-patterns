package expirations_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages/expirations"
	"github.com/stretchr/testify/assert"
)

func TestSendMessageExpirationValid(t *testing.T) {
	t.Parallel()

	// Given
	message := expirations.Message{
		Deadline: time.Now().Unix() + 1000,
		Value:    "value",
	}

	expectedMessage := expirations.Resolution{
		Processed: true,
	}

	ctx := context.TODO()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	messageChannel := make(chan expirations.Message)
	sender := expirations.NewMessageSender(messageChannel)
	receiver := expirations.NewMessageReceiver(messageChannel)

	receiver.Start(ctx)

	// When
	err := sender.Send(ctx, message)

	var closed bool

	var newMessage expirations.Resolution

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

func TestSendMessageExpirationInvalid(t *testing.T) {
	t.Parallel()

	// Given
	message := expirations.Message{
		Deadline: time.Now().Unix() - 1000,
		Value:    "value",
	}

	expectedMessage := expirations.Resolution{
		Processed: false,
	}

	ctx := context.TODO()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	messageChannel := make(chan expirations.Message)
	sender := expirations.NewMessageSender(messageChannel)
	receiver := expirations.NewMessageReceiver(messageChannel)

	receiver.Start(ctx)

	// When
	err := sender.Send(ctx, message)

	var closed bool

	var newMessage expirations.Resolution

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
