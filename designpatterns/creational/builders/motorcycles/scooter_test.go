package motorcycles_test

import (
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/builders/motorcycles"
	"github.com/stretchr/testify/assert"
)

func TestScooterBuilder(t *testing.T) {
	// Given
	expectedScooter := motorcycles.Scooter{
		Serie: "12abc",
		Engine: &motorcycles.Engine{
			CC: 125,
			HP: 12,
		},
		CentralComputer: &motorcycles.CentralComputer{
			Version: "xyz123",
			Model:   "abc",
		},
		Fairing: &motorcycles.Fairing{
			Version: "f12",
			Model:   "f12-abc-2022",
		},
	}
	engineData := motorcycles.EngineSpecification{
		CylinderCapacity: 125,
		HorsePower:       12,
	}
	scooterSerie := "12abc"
	fairingID := "f12"
	centralComputerID := "xyz123"
	builder := motorcycles.NewScooterBuilder()
	// When
	newScooter := builder.New(scooterSerie).
		WithEngine(engineData).
		WithFairing(fairingID).
		WithCentralComputer(centralComputerID).
		Build()
	// Then
	assert.Equal(t, &expectedScooter, newScooter)
}
