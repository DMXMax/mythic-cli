package game

import (
	"fmt"

	"github.com/DMXMax/mge/chart"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
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

		// Get chaos value from flag (Cobra handles flag parsing automatically)
		chaos, err := cmd.Flags().GetInt8("chaos")
		if err != nil {
			return fmt.Errorf("failed to get chaos flag: %w", err)
		}

		// Validate chaos factor range
		if chaos < chart.MIN_CHAOS || chaos > chart.MAX_CHAOS {
			return fmt.Errorf("chaos must be between %d and %d", chart.MIN_CHAOS, chart.MAX_CHAOS)
		}

		// Try to find the game in the database
		var game gdb.Game
		result := db.GamesDB.Preload("Log").Where("name = ?", name).First(&game)

		if result.Error == nil {
			// Game found, set it as current
			gdb.Current = &game
			log.Info().Str("game", name).Msg("Selected existing game")
			cmd.Printf("Selected existing game: %s (Chaos: %d)\n", name, game.Chaos)
		} else if result.Error == gorm.ErrRecordNotFound {
			// Game not found, create a new one
			newGame := &gdb.Game{
				Name:  name,
				Chaos: chaos,
			}
			// Save the new game to the database
			if err := db.GamesDB.Create(newGame).Error; err != nil {
				return fmt.Errorf("failed to create game '%s': %w", name, err)
			}
			gdb.Current = newGame
			log.Info().Str("game", name).Int8("chaos", chaos).Msg("Created new game")
			cmd.Printf("Created new game: %s (Chaos: %d)\n", name, newGame.Chaos)
		} else {
			// Another database error occurred
			return fmt.Errorf("error checking for game '%s': %w", name, result.Error)
		}

		return nil
	},
}

func init() {
	createCmd.Flags().Int8P("chaos", "x", 4, "set the chaos factor for the game")
}
