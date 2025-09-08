/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/DMXMax/mythic-cli/cmd"
	"github.com/DMXMax/mythic-cli/util/db"
	"github.com/DMXMax/mythic-cli/util/game"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	cmd.Execute()
}

func init() {
	var err error
	db.GamesDB, err = gorm.Open(sqlite.Open("data/games.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: data/games.db")
	}

	err = db.GamesDB.AutoMigrate(&game.Game{})
	if err != nil {
		panic("failed to migrate Game model")
	}
	err = db.GamesDB.AutoMigrate(&game.LogEntry{})
	if err != nil {
		panic("failed to migrate LogEntry model")
	}
}
