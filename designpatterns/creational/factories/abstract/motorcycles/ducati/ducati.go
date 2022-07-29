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

// NewFactory create a ducati factory
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
