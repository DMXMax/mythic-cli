package game

import (
	"fmt"
	"strconv"

	"github.com/DMXMax/mge/chart"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// chaosCmd sets or displays the chaos factor for the current game.
// The chaos factor (1-9) affects the likelihood of extreme results in dice rolls.
// Higher chaos values increase the chance of exceptional outcomes.
// If no value is provided, it displays the current chaos factor.
var chaosCmd = &cobra.Command{
	Use:   "chaos [value]",
	Short: "Set or show the chaos factor for the game",
	Long:  `Set or show the chaos factor for the game (1-9). Higher values increase the chance of extreme results.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g == nil {
			return fmt.Errorf("no game selected")
		}
		// if there is an argument, set the chaos value
		if len(args) == 1 {
			userChaos, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid chaos value: %s", err)
			}
			if userChaos < chart.MinChaosUser || userChaos > chart.MaxChaosUser {
				return fmt.Errorf("chaos must be between %d and %d", chart.MinChaosUser, chart.MaxChaosUser)
			}
			// Convert user input (1-9) to internal representation (0-8)
			internalChaos := int8(chart.ChaosUserToInternal(userChaos))
			g.SetChaos(internalChaos)
			fmt.Printf("Chaos factor set to %d\n", userChaos)
			// Persist the change to the database
			// Use Select() to only update chaos field, avoiding association saves
			// This prevents duplicate log entries if Log field is populated
			if err := db.GamesDB.Model(g).Select("chaos", "updated_at").Updates(map[string]interface{}{
				"chaos": internalChaos,
			}).Error; err != nil {
				return fmt.Errorf("failed to save game after changing chaos: %w", err)
			}
			return nil
		}

		// Display chaos in user-facing format (1-9)
		fmt.Printf("Current Chaos: %d\n", chart.ChaosInternalToUser(int(g.Chaos)))
		return nil
	},
}
