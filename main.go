/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/DMXMax/cli-test/cmd"
	"github.com/DMXMax/cli-test/util/db"
	"github.com/DMXMax/cli-test/util/game"
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
		panic("failed to connect database")
	}

	err = db.GamesDB.AutoMigrate(&game.Game{})
	if err != nil {
		panic("failed to migrate Game")
	}
	err = db.GamesDB.AutoMigrate(&game.LogEntry{})
	if err != nil {
		panic("failed to migrate LogEntry")
	}
}
