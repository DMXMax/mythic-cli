// Package game provides commands for managing Mythic game sessions.
// This includes creating, loading, saving, and configuring games.
package game

import (
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// GameCmd is the root command for game management operations.
// When invoked without subcommands, it displays information about the current game
// or shows usage if no game is selected.
var GameCmd = &cobra.Command{
	Use:     "game",
	Aliases: []string{"g"},
	Short:   "Manage game sessions",
	Long: `Manage Mythic game sessions including creation, loading, saving, and configuration.

Games have:
- A name for identification
- A chaos factor (1-9) that affects dice roll outcomes
- Default odds (0-8) for fate chart rolls
- A story log of all events and dice rolls
- Story themes for plot generation

If a game is currently selected, this command displays its information.
If no game is selected, it shows usage information.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g != nil {
			cmd.Println("Current Game:", *g)
		} else {
			cmd.Usage()
		}

		return nil
	},
}

func init() {
	GameCmd.AddCommand(createCmd)
	GameCmd.AddCommand(saveCmd)
	GameCmd.AddCommand(chaosCmd)
	GameCmd.AddCommand(loadCmd)
	GameCmd.AddCommand(gameListCmd)
	GameCmd.AddCommand(oddsCmd)
	GameCmd.AddCommand(removeCmd)
	GameCmd.AddCommand(exportCmd)
	GameCmd.AddCommand(infoCmd)
	GameCmd.AddCommand(plotPointCmd)
}
