/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"path/filepath"

	"github.com/DMXMax/mge/storage"
	"github.com/DMXMax/mythic-cli/cmd"
	"github.com/DMXMax/mythic-cli/util/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// main is the entry point for the Mythic CLI application.
// It initializes the database connection and executes the root command.
func main() {
	cmd.Execute()
}

// init initializes the application by setting up logging and database connections.
// It creates the data directory if it doesn't exist, opens a SQLite database connection,
// and performs automatic database migrations for Game and LogEntry models.
func init() {
	// Set logging level to Error to hide all non-critical messages from users
	log.Logger = log.Level(zerolog.ErrorLevel)

	// Use default path in home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("failed to get home directory: " + err.Error())
	}
	dbPath := filepath.Join(homeDir, ".mythic-db", "games.db")

	// Initialize database using shared storage package
	db.GamesDB, err = storage.InitDatabase(dbPath)
	if err != nil {
		panic("failed to connect database " + dbPath + ": " + err.Error())
	}

	// Run migrations for all models (including Thread/Character/Scene for future compatibility)
	err = db.GamesDB.AutoMigrate(&storage.Game{}, &storage.LogEntry{}, &storage.Thread{}, &storage.Character{}, &storage.Scene{})
	if err != nil {
		panic("failed to migrate database models: " + err.Error())
	}
}
