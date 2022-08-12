package games

import (
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/avatars"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/characters"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/motorcycles"
	"github.com/fernandoocampo/micro-patterns/designpatterns/structural/facades/games/worlds"
)

type Motorcycle interface {
	SpeedUp()
	Brake()
	Lean()
}

type Game struct {
	Name       string
	Characters []*characters.Character
	Motorcycle Motorcycle
	Avatar     *avatars.Avatar
	World      *worlds.World
}

func NewGame(name string) *Game {
	newGame := new(Game)
	newGame.Name = name

	return newGame
}

func (g *Game) Start() {
	g.Avatar = avatars.GenerateAvatar()
	g.Characters = characters.GenerateCharacters()
	g.Motorcycle = motorcycles.GenerateScooter()
	g.World = worlds.Generate()
}
