package game

import (
	"github.com/DMXMax/cli-test/util/db"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type LogEntry struct {
	gorm.Model
	Type   int
	Msg    string
	GameID uint
}

type Character struct {
	Name    string
	HighC   string
	Trouble string
	Skills  map[string]int
}
type Game struct {
	gorm.Model
	Name  string `gorm:"unique"` // Name of the game
	Chaos int8   // Current Chaos level
	Log   []LogEntry
}

type Games map[string]*Game

// Current is the current game running. Its nil if there is no game set
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

func (g *Game) AddtoGameLog(t int, msg string) {

	if g.ID == 0 { //Game has not been saved yet
		db.GamesDB.Save(&g)
	}
	entry := LogEntry{Type: t, Msg: msg, GameID: g.ID}
	g.Log = append(g.Log, entry)

	result := db.GamesDB.Save(&g)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (g *Game) GetGameLog(n int) error {

	result := db.GamesDB.Preload("Log", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at DESC").Limit(n).Find(&g.Log)
	}).Find(&g)

	return result.Error
}
