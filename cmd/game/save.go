package game

import (
	"fmt"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// saveCmd saves the current game to the database.
// This persists all game state including chaos factor and log entries.
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the current game to the database",
	Long:  `Save the current game and all its associated data (chaos factor, log entries) to the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		// Use Select() to only update game fields, avoiding association saves
		// This prevents duplicate log entries if Log field is populated
		if err := db.GamesDB.Model(g).Select("name", "chaos", "story_themes", "updated_at").Updates(g).Error; err != nil {
			return fmt.Errorf("failed to save game: %w", err)
		}
		fmt.Printf("Game '%s' saved.\n", g.Name)

		return nil
	},
}
