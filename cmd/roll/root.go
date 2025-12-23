// Package roll provides commands for rolling dice on the Mythic Fate Chart
// and rolling Fate/Fudge dice (4dF).
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

// RollCmd rolls on the Mythic Fate Chart using the current game's chaos factor.
// The chaos factor and odds can be overridden with flags.
// An optional message can be provided which will be logged with the result.
var RollCmd = &cobra.Command{
	Use:   "roll [message]",
	Short: "Roll on the Mythic Fate Chart",
	Long: `Roll on the Mythic chart using the game's chaos factor.
A message for the roll is optional. If provided, it will be logged with the result.
The chaos factor can be set with the -c flag.
The odds can be set with the -o flag (default: 50/50). Use -o ? to list all available odds.`,
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

	// Parse odds - default to 50/50 (FiftyFifty = 4) if not explicitly set
	if !cmd.Flags().Changed("odds") {
		odds = chart.FiftyFifty // Default to 50/50
	} else {
		// Normalize input for odds string
		normalized := normalizeOddsInput(oddsStr)

		// Provide helper listing when -o ? is used
		if normalized == "?" {
			printOddsHelp()
			return nil
		}

		// Try numeric odds first
		parsed, err := strconv.ParseInt(normalized, 10, 8)
		if err == nil {
			if parsed < 0 || parsed > 8 {
				return fmt.Errorf("odds must be between 0 and 8")
			}
			odds = chart.Odds(parsed)
		} else { // not a number, try to match it to a string
			matches := chart.MatchOddsPrefix(normalized)
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
		// Create log entry directly in database to avoid duplicates
		entry := gdb.LogEntry{Type: 1, Msg: logMessage, GameID: gdb.Current.ID}
		if err := db.GamesDB.Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to save log entry: %w", err)
		}
	}

	return nil

}

func init() {
	RollCmd.Flags().Int8P("chaos", "c", 4, "set the chaos factor for the game")
	RollCmd.Flags().StringP("odds", "o", "fifty", "set the odds for the roll (name or number, default: 50/50, use -o ? to list)")
	RollCmd.AddCommand(RollFateCmd)
}

// normalizeOddsInput normalizes odds input by lowercasing, trimming, and standardizing variants.
// It handles common separators (hyphens, underscores, em dashes) and converts them to spaces,
// and recognizes common numeric forms like "50/50" and converts them to "fifty fifty".
func normalizeOddsInput(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	// normalize common separators to spaces
	s = strings.ReplaceAll(s, "—", " ") // em dash
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	// collapse all whitespace sequences
	s = strings.Join(strings.Fields(s), " ")
	// common alias for fifty fifty
	switch s {
	case "50/50", "50-50", "50 50", "50\u00A050":
		return "fifty fifty"
	}
	return s
}

// printOddsHelp prints all available odds names and their numeric indices.
// This is displayed when the user runs "roll -o ?".
func printOddsHelp() {
	fmt.Println("Available odds:")
	for i, name := range chart.OddsStrList {
		fmt.Printf("%d : %s\n", i, name)
	}
}
