package messages_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages"
	"github.com/stretchr/testify/assert"
)

var errContextCancelled = errors.New("context was cancelled")

func TestRequestReply(t *testing.T) {
	t.Parallel()

	// Given
	expectedReply := messages.Reply{
		Err:     nil,
		Message: "accepted",
	}
	request := messages.Request{
		ID:   1,
		Code: "A",
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	receiver := messages.NewRequestReplyReceiver()
	go receiver.Start(ctx)

	sender := messages.NewRequestReplySender(receiver.RequestStream())
	// When
	sender.Send(ctx, request)
	reply, err := readReply(ctx, t, sender)
	// Then
	assert.NoError(t, err)
	assert.Equal(t, expectedReply, reply)
}

func readReply(ctx context.Context, t *testing.T, sender *messages.Requester) (messages.Reply, error) {
	t.Helper()

	select {
	case <-ctx.Done():
		return messages.Reply{}, errContextCancelled
	case reply, ok := <-sender.ReplyStream():
		if !ok {
			t.Fatalf("reply stream was closed unexpectedly")
		}

		return reply, nil
	}
}
