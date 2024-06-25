package db

import (
	"github.com/DMXMax/cli-test/util/game"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var GamesDB *gorm.DB

func Start() {
	var err error

	GamesDB, err = gorm.Open(sqlite.Open("data/games.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = GamesDB.AutoMigrate(&game.Game{})
	if err != nil {
		panic("failed to migrate Game")
	}
	err = GamesDB.AutoMigrate(&game.LogEntry{})
	if err != nil {
		panic("failed to migrate LogEntry")
	}

}

func init() {
	Start()
}
