package roll

import (
	"fmt"

	gdb "github.com/DMXMax/cli-test/util/game"
	"github.com/DMXMax/mge/chart"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	oddsStr string
	chaos   int8
)

// rootCmd represents the base command when called without any subcommands
var RollCmd = &cobra.Command{
	Use:   "roll",
	Short: "Rolls on the Mythic Fate Chart",
	Long: `Rolls on the Mythic chart using the game chosen chaos factor and
	provided odds. If the odds are not selected, they remain at 50/50.
	The chaos factor can be set with the -c flag.
	The odds can be set with the -o flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//	chaos = 4 // default chaos
		//	cmd.Flags().Parse(args)
		g := gdb.Current
		if g != nil {
			if !cmd.Flags().Changed("chaos") {
				chaos = g.Chaos
			}
		}

		// odds category comes from taking the odds parameter and converting it to a chart.Odds
		var odds, ok = chart.FateChartNames[oddsStr]
		if !ok {
			odds = chart.FiftyFifty
		}

		log.Trace().Msg("Rolling the dice")

		log.Trace().Str("odds", odds.String()).Int8("chaos", chaos).Msg("Rolling the dice")
		result := chart.FateChart.RollOdds(odds, int(chaos))

		fmt.Println(result)
		if gdb.Current != nil {
			gdb.Current.AddStoryEntry(1, result.String())
		}

		return nil
	},
}

func init() {
	RollCmd.Flags().Int8VarP(&chaos, "chaos", "c", 4, "set the chaos factor for the game")
	RollCmd.Flags().StringVarP(&oddsStr, "odds", "o", "50/50", "set the odds for the roll")
}
