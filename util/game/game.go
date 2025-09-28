package game

import (
	"time"

	"github.com/DMXMax/mge/util/theme"
	"github.com/DMXMax/mythic-cli/util/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogEntry struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Type      int
	Msg       string
	GameID    uuid.UUID `gorm:"type:uuid"`
}

func (l *LogEntry) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}

type Character struct {
	Name    string
	HighC   string
	Trouble string
	Skills  map[string]int
}
type Game struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         // Name of the game
	Chaos       int8           // Current Chaos level
	Odds        int8           // Current default odds
	StoryThemes theme.Themes   `gorm:"type:text"`
	Log         []LogEntry     `gorm:"foreignKey:GameID"`
}

func (g *Game) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}

// Current is the current game running. Its nil if there is no game set
var Current *Game

func (g *Game) SetChaos(v int8) {
	g.Chaos = v
}

func (g *Game) SetOdds(v int8) {
	g.Odds = v
}

func (g *Game) AddtoGameLog(t int, msg string) {
	// This function now only modifies the in-memory struct.
	// The command handler is responsible for persistence.
	entry := LogEntry{Type: t, Msg: msg}
	g.Log = append(g.Log, entry)
}

func (g *Game) GetGameLog(n int) error {

	result := db.GamesDB.Preload("Log", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at DESC").Limit(n).Find(&g.Log)
	}).Find(&g)

	return result.Error
}
