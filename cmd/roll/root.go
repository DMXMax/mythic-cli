package roll

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/DMXMax/mge/chart"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RollCmd = &cobra.Command{
	Use:   "roll [message]",
	Short: "Rolls on the Mythic Fate Chart",
	Long: `Rolls on the Mythic chart using the game's chaos factor and odds.
A message for the roll is optional. If provided, it will be logged with the result.
	The chaos factor can be set with the -c flag.
	The odds can be set with the -o flag.`,
	RunE: RollFunc,
}

func RollFunc(cmd *cobra.Command, args []string) error {
	var odds chart.Odds
	messageArgs := args
	g := gdb.Current

	// Get chaos value from flag
	chaosValue, err := cmd.Flags().GetInt8("chaos")
	if err != nil {
		return fmt.Errorf("failed to get chaos flag: %w", err)
	}

	// Get odds value from flag
	oddsStr, err := cmd.Flags().GetString("odds")
	if err != nil {
		return fmt.Errorf("failed to get odds flag: %w", err)
	}

	// Use game's chaos if not explicitly set, otherwise use default
	if !cmd.Flags().Changed("chaos") {
		if g != nil {
			chaosValue = g.Chaos
		} else {
			chaosValue = 4 // default chaos value
		}
	}

	// Use game's odds if not explicitly set, otherwise parse the flag
	if !cmd.Flags().Changed("odds") && g != nil {
		odds = chart.Odds(g.Odds)
	} else {
		parsed, err := strconv.ParseInt(oddsStr, 10, 8)
		if err == nil {
			if parsed < 0 || parsed > 8 {
				return fmt.Errorf("odds must be between 0 and 8")
			}
			odds = chart.Odds(parsed)
		} else { // not a number, try to match it to a string
			// Try to combine with next arg for multi-word odds
			if len(messageArgs) > 0 {
				combined := oddsStr + " " + messageArgs[0]
				// If combined is a more specific match, use it
				originalMatches := chart.MatchOddsPrefix(oddsStr)
				newMatches := chart.MatchOddsPrefix(combined)
				if len(newMatches) > 0 && (len(originalMatches) > len(newMatches) || len(originalMatches) == 0) {
					oddsStr = combined
					messageArgs = messageArgs[1:]
				} else if len(originalMatches) == 1 && len(newMatches) == 1 {
					// Handle "fifty" vs "fifty fifty" where both might give a single match
					if strings.EqualFold(chart.OddsStrList[newMatches[0]], combined) {
						oddsStr = combined
						messageArgs = messageArgs[1:]
					}
				}
			}

			matches := chart.MatchOddsPrefix(oddsStr)
			if len(matches) == 0 {
				err := fmt.Errorf("invalid odds: '%s'", oddsStr)
				log.Error().Err(err).Msg("Invalid odds")
				return err

			}
			if len(matches) != 1 { // multiple possible odds
				fmt.Println("Did you mean one of these odds?")
				for _, match := range matches {
					fmt.Printf("%d : %s\n", match, chart.OddsStrList[match])
				}
				return fmt.Errorf("multiple possible odds for '%s'", oddsStr)
			}

			odds = chart.Odds(matches[0])
		}
	}

	message := strings.Join(messageArgs, " ")
	if len(message) > 256 {
		return fmt.Errorf("message cannot be longer than 256 characters")
	}
	result := chart.FateChart.RollOdds(odds, int(chaosValue))

	logMessage := strings.TrimSpace(fmt.Sprintf("%s (C:%d) -> %s", message, chaosValue, result))

	fmt.Println(logMessage)
	if gdb.Current != nil {
		gdb.Current.AddtoGameLog(1, logMessage)
		// Persist the game state, including the new log entry
		if err := db.GamesDB.Save(gdb.Current).Error; err != nil {
			return fmt.Errorf("failed to save game after roll: %w", err)
		}
	}

	return nil

}

func init() {
	RollCmd.Flags().Int8P("chaos", "c", 4, "set the chaos factor for the game")
	RollCmd.Flags().StringP("odds", "o", "5", "set the odds for the roll (name or number)")
}
