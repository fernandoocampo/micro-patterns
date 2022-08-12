package requestreplies_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages/requestreplies"
	"github.com/stretchr/testify/assert"
)

var errContextCancelled = errors.New("context was cancelled")

func TestRequestReply(t *testing.T) {
	t.Parallel()

	// Given
	expectedReply := requestreplies.Reply{
		Err:     nil,
		Message: "accepted",
	}
	request := requestreplies.Request{
		ID:   1,
		Code: "A",
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	receiver := requestreplies.NewRequestReplyReceiver()
	go receiver.Start(ctx)

	sender := requestreplies.NewRequestReplySender(receiver.RequestStream())
	// When
	sender.Send(ctx, request)
	reply, err := readReply(ctx, t, sender)
	// Then
	assert.NoError(t, err)
	assert.Equal(t, expectedReply, reply)
}

func readReply(ctx context.Context, t *testing.T, sender *requestreplies.Requester) (requestreplies.Reply, error) {
	t.Helper()

	select {
	case <-ctx.Done():
		return requestreplies.Reply{}, errContextCancelled
	case reply, ok := <-sender.ReplyStream():
		if !ok {
			t.Fatalf("reply stream was closed unexpectedly")
		}

		return reply, nil
	}
}
