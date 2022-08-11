package factories

import (
	"github.com/fernandoocampo/micro-patterns/designpatterns/creational/factories/abstract/motorcycles"
)

// factory defines behavior to create urban motorcycle products.
type UrbanCreator interface {
	CreateUrban() motorcycles.UrbanBehavior
}

// factory defines behavior to create sport motorcycle products.
type SportCreator interface {
	CreateSport() motorcycles.SportBehavior
}

// factory defines behavior to create adventure motorcycle products.
type AdventureCreator interface {
	CreateAdventure() motorcycles.AdventureBehavior
}
