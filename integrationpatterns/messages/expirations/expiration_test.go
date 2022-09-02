package expirations_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages/expirations"
	"github.com/stretchr/testify/assert"
)

func TestSendMessageExpiration(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		message expirations.Message
		want    expirations.Resolution
	}{
		"valid": {
			message: expirations.Message{
				Deadline: time.Now().Unix() + 1000,
				Value:    "value",
			},
			want: expirations.Resolution{
				Processed: true,
			},
		},
		"invalid": {
			message: expirations.Message{
				Deadline: time.Now().Unix() - 1000,
				Value:    "value",
			},
			want: expirations.Resolution{
				Processed: false,
			},
		},
	}

	for name, data := range cases {
		name, data := name, data

		t.Run(name, func(st *testing.T) {
			st.Parallel()

			newMessage, err := send(st, data.message)
			// Then
			assert.NoError(t, err)
			assert.Equal(t, data.want, newMessage)
		})
	}
}

func send(t *testing.T, message expirations.Message) (expirations.Resolution, error) {
	t.Helper()

	// Given
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	messageChannel := make(chan expirations.Message)
	sender := expirations.NewMessageSender(messageChannel)
	receiver := expirations.NewMessageReceiver(messageChannel)

	receiver.Start(ctx)

	// When
	err := sender.Send(ctx, message)
	if err != nil {
		return expirations.Resolution{}, fmt.Errorf("failed: %w", err)
	}

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

	return newMessage, nil
}
