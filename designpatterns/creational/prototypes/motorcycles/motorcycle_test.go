package motorcycles_test

import (
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/prototypes/motorcycles"
	"github.com/stretchr/testify/assert"
)

func TestCloneMotorcycle(t *testing.T) {
	// Given
	motorcycle := &motorcycles.Motorcycle{
		Brand: "bmw",
		Name:  "nymeria",
		Model: 2019,
		Miles: 23000,
	}
	// WHEN
	clonedMotorcycle := callClone(t, motorcycle)
	// THEN
	assert.NotEqual(t, clonedMotorcycle, &motorcycle)
	assert.Equal(t, *motorcycle, *clonedMotorcycle)
}

type CloneableMotorcycles interface {
	Clone() motorcycles.Motorcycle
}

func callClone(t *testing.T, m CloneableMotorcycles) *motorcycles.Motorcycle {
	t.Helper()
	clonedMotorcycle := m.Clone()
	return &clonedMotorcycle
}
