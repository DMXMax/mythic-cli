package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/DMXMax/mge/chart"
	"github.com/DMXMax/mge/storage"
	"github.com/DMXMax/mge/util/theme"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// createCmd creates a new game or selects an existing one if a game with the same name exists.
// If a game with the given name already exists, it is loaded and set as the current game.
// Otherwise, a new game is created with the specified chaos factor (default: 4).
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "new"},
	Short:   "Create a new game or select an existing one",
	Long: `Create a new game with a supplied name. If a game with that name already exists,
it will be selected and set as the current game instead of creating a duplicate.

The chaos factor can be specified with the --chaos or -x flag (default: 4).
Valid chaos factor range is 1-9.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required arguments
		if len(args) < 1 {
			return fmt.Errorf("create requires a name")
		}

		// Join all args to handle multi-word names (e.g., "Kat in Shadow")
		name := storage.SanitizeGameName(strings.Join(args, " "))

		// Validate game name using shared validation
		if err := storage.ValidateGameName(name); err != nil {
			return err
		}

		// Get chaos value from flag (Cobra handles flag parsing automatically)
		chaos, err := cmd.Flags().GetInt8("chaos")
		if err != nil {
			return fmt.Errorf("failed to get chaos flag: %w", err)
		}

		// Validate chaos factor range
		if chaos < chart.MinChaos || chaos > chart.MaxChaos {
			return fmt.Errorf("chaos must be between %d and %d", chart.MinChaos, chart.MaxChaos)
		}

		// Try to find the game in the database
		var game gdb.Game
		result := db.GamesDB.Where("name = ?", name).First(&game)

		switch err := result.Error; {
		case err == nil:
			// Game found, set it as current
			gdb.Current = &game
			log.Info().Str("game", name).Msg("Selected existing game")
			cmd.Printf("Game '%s' already exists - loaded existing game (Chaos: %d)\n", name, game.Chaos)
		case errors.Is(err, gorm.ErrRecordNotFound):
			// Game not found, create a new one

			newGame := &gdb.Game{
				Name:        name,
				Chaos:       chaos,
				StoryThemes: theme.GetThemes(),
			}
			// Save the new game to the database
			if err := db.GamesDB.Create(newGame).Error; err != nil {
				return fmt.Errorf("failed to create game '%s': %w", name, err)
			}
			gdb.Current = newGame
			log.Info().Str("game", name).Int8("chaos", chaos).Msg("Created new game")
			cmd.Printf("Created new game: %s (Chaos: %d)\n", name, newGame.Chaos)
		default:
			// Another database error occurred
			return fmt.Errorf("error checking for game '%s': %w", name, err)
		}

		return nil
	},
}

func init() {
	createCmd.Flags().Int8P("chaos", "x", 4, "set the chaos factor for the game")
}
