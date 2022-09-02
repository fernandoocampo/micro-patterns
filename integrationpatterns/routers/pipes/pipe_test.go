package pipes_test

import (
	"context"
	"testing"
	"time"

	"github.com/fernandoocampo/micro-patterns/integrationpatterns/routers/pipes"
	"github.com/stretchr/testify/assert"
)

func TestProcessEvent(t *testing.T) {
	t.Parallel()
	// Given
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	loggerChannel := pipes.NewPipe()
	storageChannel := pipes.NewPipe()
	enrichChannel := pipes.NewPipe()
	validateChannel := pipes.NewPipe()
	collectorChannel := pipes.NewPipe()
	endChannel := pipes.NewPipe()

	storage := pipes.NewStorage(storageChannel).
		WithNextFilter(endChannel)
	storage.Start(ctx)

	logger := pipes.NewLogger(loggerChannel).
		WithNextFilter(storageChannel)
	logger.Start(ctx)

	enricher := pipes.NewEnrich(enrichChannel).
		WithNextFilter(loggerChannel)
	enricher.Start(ctx)

	validator := pipes.NewValidator(validateChannel).
		WithNextFilter(enrichChannel)
	validator.Start(ctx)

	collector := pipes.NewCollector(collectorChannel).
		WithNextFilter(validateChannel)
	collector.Start(ctx)

	expectedStoredValue := pipes.Order{
		ID:         "1A",
		InternalID: "1A-Z1",
		Value:      7.32,
		Success:    true,
	}
	expectedLog := `{"id":"1A","internal_id":"1A-Z1","value":7.32}`
	expectedValidation := true
	expectedEnrichment := "1A-Z1"
	expectedCollectedValue := pipes.NewOrder{
		ID:    "1A",
		Value: 7.32,
	}

	newEvent := pipes.NewOrder{
		ID:    "1A",
		Value: 7.32,
	}
	// When
	collectorChannel <- newEvent

	<-endChannel
	// Then
	assert.Equal(t, expectedCollectedValue, collector.Value())
	assert.Equal(t, expectedValidation, validator.Result())
	assert.Equal(t, expectedEnrichment, enricher.GetGeneratedID())
	assert.Equal(t, expectedLog, logger.Logs())
	assert.Equal(t, expectedStoredValue, storage.GetOrder())
}
