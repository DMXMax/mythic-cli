// Package game provides the core game data structures and operations
// for the Mythic CLI application, including Game and LogEntry models.
package game

import (
	"time"

	"github.com/DMXMax/mge/util/theme"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LogEntry represents a single entry in a game's story log.
// Log entries can be of different types (e.g., dice rolls, story events)
// and are automatically timestamped.
type LogEntry struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time      // When the log entry was created
	UpdatedAt time.Time      // When the log entry was last updated
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete support
	Type      int            // Type of log entry (0 = story, 1 = dice roll, etc.)
	Msg       string         // The log message content
	GameID    uuid.UUID      `gorm:"type:uuid"` // Foreign key to the game
}

// BeforeCreate is a GORM hook that generates a UUID for the log entry before creation.
func (l *LogEntry) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}

// Character represents a character in the game.
// This structure is defined but not currently used in the database schema.
type Character struct {
	Name    string         // Character name
	HighC   string         // High concept
	Trouble string         // Character trouble/aspect
	Skills  map[string]int // Map of skill names to values
}

// Game represents a Mythic game session with all its associated data.
// Each game has a name, chaos factor, story themes, and a log of events.
type Game struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt   time.Time      // When the game was created
	UpdatedAt   time.Time      // When the game was last updated
	DeletedAt   gorm.DeletedAt `gorm:"index"` // Soft delete support
	Name        string         `gorm:"uniqueIndex"` // Name of the game (unique)
	Chaos       int8           // Current Chaos level (1-9)
	StoryThemes theme.Themes   `gorm:"type:text"`         // Story themes for plot generation
	Log         []LogEntry     `gorm:"foreignKey:GameID"` // Log entries (deprecated: only used for export template; query database directly for log entries)
}

// BeforeCreate is a GORM hook that generates a UUID for the game before creation.
func (g *Game) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}

// Current is the currently active game session.
// It is nil if no game has been loaded or created.
var Current *Game

// SetChaos sets the chaos factor for the game.
// The chaos factor affects the likelihood of extreme results in dice rolls.
// Valid range is 1-9.
func (g *Game) SetChaos(v int8) {
	g.Chaos = v
}

