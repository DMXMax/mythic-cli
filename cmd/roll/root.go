package roll

import (
	"fmt"

	"strconv"

	"github.com/DMXMax/mge/chart"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	OddsStrList string
	chaos       int8
)

// rootCmd represents the base command when called without any subcommands
var RollCmd = &cobra.Command{
	Use:   "roll",
	Short: "Rolls on the Mythic Fate Chart",
	Long: `Rolls on the Mythic chart using the game chosen chaos factor and
	provided odds. If the odds are not selected, they remain at 50/50.
	The chaos factor can be set with the -c flag.
	The odds can be set with the -o flag.`,
	RunE: RollFunc,
}

// if there's a game, use its chaos value. If not, use a default of 5gl unless its set.
func RollFunc(cmd *cobra.Command, args []string) error {
	var odds chart.Odds
	g := gdb.Current
	if g != nil {
		if !cmd.Flags().Changed("chaos") {
			chaos = g.Chaos
		}
	} else {
		// No active game
		if !cmd.Flags().Changed("chaos") {
			chaos = 4
		}

		if !cmd.Flags().Changed("odds") {
			odds = 5
		}
	}

	// try to convert odds to a number. If not, try to match it to a string
	parsed, err := strconv.ParseInt(OddsStrList, 10, 8)
	if err != nil { //not a number, try to match it to a string
		matches := chart.MatchOddNametoOdds(OddsStrList)
		if len(matches) == 0 {
			err := fmt.Errorf("invalid odds: %s", OddsStrList)
			log.Error().Err(err).Msg("Invalid odds")
			return err

		}
		if len(matches) != 1 { //multiple possible odds

			fmt.Println("Did you mean one of these odds?")
			for _, match := range matches {
				fmt.Printf("%d : %s\n", match, chart.OddsStrList[match])
			}
			return fmt.Errorf("multiple possible odds")
		}

		odds = chart.Odds(matches[0])

	} else {
		odds = chart.Odds(parsed)
	}

	log.Trace().Str("odds", odds.String()).Int8("chaos", chaos).Msg("Rolling the dice")
	result := chart.FateChart.RollOdds(odds, int(chaos))

	fmt.Println(result)
	if gdb.Current != nil {
		gdb.Current.AddtoGameLog(1, result.String())
	}

	return nil

}

func init() {
	RollCmd.Flags().Int8VarP(&chaos, "chaos", "c", 4, "set the chaos factor for the game")
	RollCmd.Flags().StringVarP(&OddsStrList, "odds", "o", "4", "set the odds for the roll")
}
