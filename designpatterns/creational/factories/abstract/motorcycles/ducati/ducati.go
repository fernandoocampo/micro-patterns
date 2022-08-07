package ducati

import (
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/adventures"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/sports"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati/urbans"
)

// Factory creates  motorcycles.
type Factory struct {
}

type UrbanFactory[T any] struct {
	MotorcycleType T
}

type SportFactory[T any] struct {
	MotorcycleType T
}

type AdventureFactory[T any] struct {
	MotorcycleType T
}

// NewFactory create a bmw factory
func NewFactory() *Factory {
	newFactory := Factory{}

	return &newFactory
}

func (f *Factory) CreateUrban() motorcycles.UrbanBehavior {
	return urbans.New()
}
func (f *UrbanFactory[T]) CreateUrban() T {
	var bike interface{} = urbans.New()

	newBike, ok := bike.(T)
	if !ok {
		return newBike
	}

	return newBike
}
func (f *Factory) CreateSport() motorcycles.SportBehavior {
	return sports.New()
}
func (f *SportFactory[T]) CreateSport() T {
	var bike interface{} = sports.New()

	newBike, ok := bike.(T)
	if !ok {
		return newBike
	}

	return newBike
}
func (f *Factory) CreateAdventure() motorcycles.AdventureBehavior {
	return adventures.New()
}
func (f *AdventureFactory[T]) CreateAdventure() T {
	var bike interface{} = adventures.New()

	newBike, ok := bike.(T)
	if !ok {
		return newBike
	}

	return newBike
}
