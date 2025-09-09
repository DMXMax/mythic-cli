package game

import (
	"fmt"
	"strconv"

	"github.com/DMXMax/mge/chart"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var chaosCmd = &cobra.Command{
	Use:   "chaos [value]",
	Short: "Set or show the chaos factor for the game",
	Long:  `Set or show the chaos factor for the game (1-9).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g == nil {
			return fmt.Errorf("no game selected")
		}
		// if there is an argument, set the chaos value
		if len(args) == 1 {
			set, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid chaos value: %s", err)
			}
			if set < chart.MinChaos || set > chart.MaxChaos {
				return fmt.Errorf("chaos must be between %d and %d", chart.MinChaos, chart.MaxChaos)
			}
			g.SetChaos(int8(set))
			fmt.Printf("Chaos factor set to %d\n", set)
			// Persist the change to the database
			if err := db.GamesDB.Save(g).Error; err != nil {
				return fmt.Errorf("failed to save game after changing chaos: %w", err)
			}
			return nil
		}

		fmt.Printf("Current Chaos: %d\n", g.Chaos)
		return nil
	},
}
