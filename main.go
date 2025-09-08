/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/DMXMax/mythic-cli/cmd"
	"github.com/DMXMax/mythic-cli/util/db"
	"github.com/DMXMax/mythic-cli/util/game"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	cmd.Execute()
}

func init() {
	// Set logging level to Error to hide all non-critical messages from users
	log.Logger = log.Level(zerolog.ErrorLevel)

	var err error
	db.GamesDB, err = gorm.Open(sqlite.Open("data/games.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database: data/games.db")
	}

	err = db.GamesDB.AutoMigrate(&game.Game{}, &game.LogEntry{})
	if err != nil {
		panic("failed to migrate database models")
	}
}
