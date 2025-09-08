package game

import (
	"fmt"

	"github.com/DMXMax/mge/chart"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// createCmd represents the command to create or select a game
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "new"},
	Short:   "create or select a game",
	Long:    `Create a new game with a supplied name. If the game already exists, it will be selected.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required arguments
		if len(args) < 1 {
			return fmt.Errorf("create requires a name")
		}

		name := args[0]

		// Parse flags before accessing them
		if err := cmd.ParseFlags(args); err != nil {
			return fmt.Errorf("failed to parse flags: %w", err)
		}

		// Get chaos value from flag
		chaos, err := cmd.Flags().GetInt8("chaos")
		if err != nil {
			return fmt.Errorf("failed to get chaos flag: %w", err)
		}

		// Validate chaos factor range
		if chaos < chart.MIN_CHAOS || chaos > chart.MAX_CHAOS {
			return fmt.Errorf("chaos must be between %d and %d", chart.MIN_CHAOS, chart.MAX_CHAOS)
		}

		// Check if game already exists
		existingGame := gdb.GetGame(name)
		if existingGame != nil {
			gdb.Current = existingGame
			log.Info().Str("game", name).Msg("Selected existing game")
			cmd.Printf("Selected existing game: %s (Chaos: %d)\n", name, existingGame.Chaos)
		} else {
			// Create new game
			game := gdb.SetGame(name, chaos)
			log.Info().Str("game", name).Int8("chaos", chaos).Msg("Created new game")
			cmd.Printf("Created new game: %s (Chaos: %d)\n", name, game.Chaos)
		}

		return nil
	},
}

func init() {
	createCmd.Flags().Int8P("chaos", "x", 4, "set the chaos factor for the game")
}
