package game

import (
	"github.com/rs/zerolog/log"
)

type Game struct {
	Name       string
	Properties map[string]any
}
type Games map[string]*Game

var Current *Game

var games Games

func SetGame(name string) *Game {
	//games is a singleton
	if games == nil {
		games = make(Games)
	}
	if g, ok := games[name]; ok {
		Current = g
		log.Info().Msgf("Game already exists: %s", name)
		return g
	}
	g := Game{}
	g.Name = name
	//games[name] = &g

	Current = &g

	log.Info().Msgf("New game created: %s", name)

	return &g

}

func GetGame(name string) *Game {

	return games[name]

}
