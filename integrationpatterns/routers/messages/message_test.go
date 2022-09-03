package messages_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/routers/messages"
	"github.com/stretchr/testify/assert"
)

func TestProcessEvent(t *testing.T) {
	t.Parallel()
	// Given
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	pipeline := newStoragePipe()
	startChannel, endChannel := pipeline.build(ctx)
	expectedStoredLog := "order id or order value is invalid"
	expectedStoredValue := messages.Order{
		ID:         "1A",
		InternalID: "1A-Z1",
		Value:      7.32,
		Success:    true,
	}
	expectedLogs := []string{
		"order id or order value is invalid",
	}
	expectedValidation := []bool{true, false}
	expectedEnrichment := "1A-Z1"
	expectedCollectedValues := []messages.NewOrder{
		{
			ID:    "1A",
			Value: 7.32,
		},
		{
			ID:    "2A",
			Value: -2.34,
		},
	}

	newValidEvent := messages.NewOrder{
		ID:    "1A",
		Value: 7.32,
	}
	newInvalidEvent := messages.NewOrder{
		ID:    "2A",
		Value: -2.34,
	}
	// When
	startChannel <- newValidEvent
	startChannel <- newInvalidEvent

	<-endChannel
	<-endChannel
	// Then
	assert.Equal(t, expectedCollectedValues, pipeline.collector.Values())
	assert.Equal(t, expectedValidation, pipeline.validator.Result())
	assert.Equal(t, expectedEnrichment, pipeline.enricher.GetGeneratedID())
	assert.Equal(t, expectedLogs, pipeline.logger.Logs())
	assert.Equal(t, expectedStoredLog, pipeline.storage.GetLog())
	assert.Equal(t, expectedStoredValue, pipeline.storage.GetOrder())
}

type storagePipe struct {
	storage   *messages.Storage
	logger    *messages.Logger
	enricher  *messages.Enrich
	validator *messages.Validator
	collector *messages.Collector
}

func newStoragePipe() *storagePipe {
	return new(storagePipe)
}

func (s *storagePipe) build(ctx context.Context) (chan<- interface{}, <-chan interface{}) {
	loggerChannel := messages.NewPipe()
	storageChannel := messages.NewPipe()
	enrichChannel := messages.NewPipe()
	validateChannel := messages.NewPipe()
	collectorChannel := messages.NewPipe()
	endChannel := messages.NewPipe()

	s.storage = messages.NewStorage(storageChannel).
		WithNextFilter(endChannel)
	s.storage.Start(ctx)

	s.logger = messages.NewLogger(loggerChannel).
		WithNextFilter(storageChannel)
	s.logger.Start(ctx)

	s.enricher = messages.NewEnrich(enrichChannel).
		WithNextFilter(storageChannel)
	s.enricher.Start(ctx)

	s.validator = messages.NewValidator(validateChannel).
		WithEnrichFilter(enrichChannel).
		WithLoggerFilter(loggerChannel)
	s.validator.Start(ctx)

	s.collector = messages.NewCollector(collectorChannel).
		WithNextFilter(validateChannel)
	s.collector.Start(ctx)

	return collectorChannel, endChannel
}
