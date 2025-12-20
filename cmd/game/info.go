package game

import (
	"fmt"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// infoCmd displays comprehensive information about the current game,
// including its name, story themes, and the most recent log entries.
var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"i"},
	Short:   "Display information about the current game",
	Long:    `Displays the name, themes, and the last 5 log entries for the currently selected game.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g == nil {
			return fmt.Errorf("no game selected. Use 'game load <name>' to select one")
		}

		cmd.Printf("Game: %s\n", g.Name)
		cmd.Println("Themes:")
		for _, theme := range g.StoryThemes {
			cmd.Printf("- %s\n", theme.String())
		}

		cmd.Println("\nRecent Log Entries:")

		// Fetch the last 5 log entries from the database
		var entries []gdb.LogEntry
		q := db.GamesDB.Model(&gdb.LogEntry{}).
			Where("game_id = ?", g.ID).
			Order("created_at DESC").
			Limit(5)
		if err := q.Find(&entries).Error; err != nil {
			return fmt.Errorf("failed to load log entries: %w", err)
		}

		if len(entries) == 0 {
			cmd.Println("  No log entries found.")
		} else {
			// Print oldest-first for natural reading by reversing the slice
			for i := len(entries) - 1; i >= 0; i-- {
				cmd.Printf("- %s\n", entries[i].Msg)
			}
		}
		return nil
	},
}
