package game

import (
	"fmt"
	"strconv"

	"github.com/DMXMax/mge/chart"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var chaosCmd = &cobra.Command{
	Use:   "chaos",
	Short: "set the chaos factor for the game",
	Long:  `Set the chaos factor for the game.`,
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
			if set < chart.MIN_CHAOS || set > chart.MAX_CHAOS {
				return fmt.Errorf("chaos must be between %d and %d", chart.MIN_CHAOS, chart.MAX_CHAOS)
			}
			g.Chaos = int8(set)
			log.Trace().Int8("chaos", int8(set)).Msg("Chaos set")
			return nil
		}

		fmt.Printf("Current Chaos: %d\n", g.Chaos)
		return nil
	},
}
