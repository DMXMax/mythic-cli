package game

import (
    "fmt"
    "strings"

    "github.com/DMXMax/mythic-cli/util/db"
    gdb "github.com/DMXMax/mythic-cli/util/game"
    "github.com/spf13/cobra"
)

var removeName string

// removeCmd removes a game and its associated log entries
var removeCmd = &cobra.Command{
    Use:     "remove [name]",
    Aliases: []string{"rm", "delete", "del"},
    Short:   "remove a game and its logs",
    Long:    `Remove a game by name. This also removes all associated log entries. You can pass the name as a positional argument or via --name.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        var name string
        if len(args) > 0 {
            name = args[0]
        } else {
            name = removeName
        }
        name = strings.TrimSpace(name)
        if name == "" {
            return fmt.Errorf("no game name specified")
        }

        // Find the game by name
        var game gdb.Game
        if err := db.GamesDB.Where("name = ?", name).First(&game).Error; err != nil {
            return fmt.Errorf("could not find game '%s': %w", name, err)
        }

        // Delete associated log entries first
        res := db.GamesDB.Where("game_id = ?", game.ID).Delete(&gdb.LogEntry{})
        if res.Error != nil {
            return fmt.Errorf("failed to delete log entries for '%s': %w", name, res.Error)
        }
        logsRemoved := res.RowsAffected

        // Delete the game
        if err := db.GamesDB.Delete(&game).Error; err != nil {
            return fmt.Errorf("failed to delete game '%s': %w", name, err)
        }

        // Clear current selection if it was this game
        if gdb.Current != nil && gdb.Current.ID == game.ID {
            gdb.Current = nil
        }

        cmd.Printf("Removed game: %s (deleted %d log entries)\n", name, logsRemoved)
        return nil
    },
}

func init() {
    removeCmd.Flags().StringVar(&removeName, "name", "", "name of the game to remove")
}

