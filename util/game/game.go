// Package game provides application-specific game state management
// for the Mythic CLI application. Data models are imported from mge/storage.
package game

import (
	"github.com/DMXMax/mge/storage"
)

// Re-export types from storage package for convenience
type (
	Game     = storage.Game
	LogEntry = storage.LogEntry
)

// Current is the currently active game session.
// It is nil if no game has been loaded or created.
var Current *Game

