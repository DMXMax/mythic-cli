package game

import (
	"fmt"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// saveCmd saves the current game to the database.
// This persists all game state including chaos factor, odds, and log entries.
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the current game to the database",
	Long:  `Save the current game and all its associated data (chaos factor, odds, log entries) to the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		if err := db.GamesDB.Save(g).Error; err != nil {
			return fmt.Errorf("failed to save game: %w", err)
		}
		fmt.Printf("Game '%s' saved.\n", g.Name)

		return nil
	},
}
