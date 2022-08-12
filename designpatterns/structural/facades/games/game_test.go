package games_test

import (
	"testing"

	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/avatars"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/characters"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/motorcycles"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/worlds"
	"github.com/stretchr/testify/assert"
)

func TestStartNewWorld(t *testing.T) {
	t.Parallel()

	// Given
	expectedNewGame := games.Game{
		Name:       "myworld",
		Characters: charactersFixture(),
		Motorcycle: motorcycleFixture(),
		Avatar:     avatarFixture(),
		World:      worldFixture(),
	}
	gameName := "myworld"
	newGame := games.NewGame(gameName)
	// When
	newGame.Start()
	// Then
	assert.Equal(t, &expectedNewGame, newGame)
}

func charactersFixture() []*characters.Character {
	return []*characters.Character{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
		{
			ID: 3,
		},
	}
}

func motorcycleFixture() *motorcycles.Scooter {
	return &motorcycles.Scooter{
		ID:     1,
		HP:     12.0,
		Torque: 10.0,
	}
}

func avatarFixture() *avatars.Avatar {
	return &avatars.Avatar{
		Speed:  80,
		Curves: 67,
		Brakes: 20,
	}
}

func worldFixture() *worlds.World {
	return &worlds.World{
		Seed: 10000000,
	}
}
