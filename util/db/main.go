// Package db provides database utilities for the Mythic CLI application.
// It manages the global database connection used throughout the application.
package db

import (
	"gorm.io/gorm"
)

// GamesDB is the global database connection instance.
// It is initialized in main.init() and used by all database operations.
var GamesDB *gorm.DB
