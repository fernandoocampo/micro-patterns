package correlations_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/messages/correlations"
	"github.com/stretchr/testify/assert"
)

var errContextCancelled = errors.New("context was cancelled")

func TestCorrelationID(t *testing.T) {
	t.Parallel()

	// Given
	correlationID := "1"
	expectedStatus := []correlations.RequestStatus{
		{Message: "accepted"},
		{Message: "processed"},
	}
	request := correlations.Request{
		CorrelationID: correlationID,
		Code:          "A",
	}

	replyChannel := make(chan correlations.Request)
	requestChannel := make(chan correlations.Reply)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	receiver := correlations.NewRequestReplyReceiver(replyChannel, requestChannel)
	go receiver.Start(ctx)

	sender := correlations.NewRequestReplySender(replyChannel, requestChannel)
	go sender.Start(ctx)
	// When
	sender.Send(ctx, request)
	status, err := readStatus(ctx, t, sender)
	// Then
	assert.NoError(t, err)
	assert.Equal(t, expectedStatus, status)
}

func readStatus(ctx context.Context, t *testing.T, sender *correlations.Requester) ([]correlations.RequestStatus, error) {
	t.Helper()

	var result []correlations.RequestStatus

	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return result, errContextCancelled
		case reply, ok := <-sender.Audit():
			if !ok {
				t.Fatalf("reply stream was closed unexpectedly")
			}

			result = append(result, reply)
		}
	}

	return result, nil
}
