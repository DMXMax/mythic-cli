package game

import (
	"github.com/rs/zerolog/log"
)

type LogEntry struct {
	Type int
	Msg  string
}

type Character struct {
	Name    string
	HighC   string
	Trouble string
	Skills  map[string]int
}
type Game struct {
	Name       string // Name of the game
	Chaos      int8   // Current Chaos level
	Story      []LogEntry
	Properties map[string]any
}

type Games map[string]*Game

var Current *Game

var games Games

func SetGame(name string, chaos int8) *Game {
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
	g.Chaos = chaos
	//games[name] = &g

	Current = &g

	log.Info().Msgf("New game created: %s", name)

	return &g

}

func (g *Game) SetChaos(v int8) {
	g.Chaos = v
}

func GetGame(name string) *Game {

	return games[name]

}

func (g *Game) AddStoryEntry(t int, msg string) {
	g.Story = append(g.Story, LogEntry{Type: t, Msg: msg})
	log.Info().Msgf("Story Entry added: %s", msg)
}

func (g *Game) GetStory() []LogEntry {
	return g.Story
}
