package bmw

import (
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/adventures"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/sports"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw/urbans"
)

// Factory creates  motorcycles.
type Factory struct {
}

// NewFactory create a bmw factory
func NewFactory() *Factory {
	newFactory := Factory{}
	return &newFactory
}

func (f *Factory) CreateUrban() motorcycles.UrbanBehavior {
	return urbans.New()
}
func (f *Factory) CreateSport() motorcycles.SportBehavior {
	return sports.New()
}
func (f *Factory) CreateAdventure() motorcycles.AdventureBehavior {
	return adventures.New()
}
