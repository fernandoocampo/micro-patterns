package factories

import (
	"errors"

	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/bmw"
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles/ducati"
)

// factory defines behavior to create urban motorcycle products
type UrbanCreator interface {
	CreateUrban() motorcycles.UrbanBehavior
}

// factory defines behavior to create sport motorcycle products
type SportCreator interface {
	CreateSport() motorcycles.SportBehavior
}

// factory defines behavior to create adventure motorcycle products
type AdventureCreator interface {
	CreateAdventure() motorcycles.AdventureBehavior
}

// Supported motorcycle brands
const (
	BMW motorcycles.Brand = iota
	Ducati
)

var ErrUnknowBrand = errors.New("unkown brand")

// NewUrbanFactory creates an urban motorcycle factory.
func NewUrbanFactory(brand motorcycles.Brand) (UrbanCreator, error) {
	switch brand {
	case BMW:
		return &bmw.UrbanFactory[motorcycles.UrbanBehavior]{}, nil
	case Ducati:
		return &ducati.UrbanFactory[motorcycles.UrbanBehavior]{}, nil
	default:
		return nil, ErrUnknowBrand
	}
}

// NewSportFactory creates an urban motorcycle factory.
func NewSportFactory(brand motorcycles.Brand) (SportCreator, error) {
	switch brand {
	case BMW:
		return &bmw.SportFactory[motorcycles.SportBehavior]{}, nil
	case Ducati:
		return &ducati.SportFactory[motorcycles.SportBehavior]{}, nil
	default:
		return nil, ErrUnknowBrand
	}
}

// NewAdventureFactory creates an adventure motorcycle factory.
func NewAdventureFactory(brand motorcycles.Brand) (AdventureCreator, error) {
	switch brand {
	case BMW:
		return &bmw.AdventureFactory[motorcycles.AdventureBehavior]{}, nil
	case Ducati:
		return &ducati.AdventureFactory[motorcycles.AdventureBehavior]{}, nil
	default:
		return nil, ErrUnknowBrand
	}
}
