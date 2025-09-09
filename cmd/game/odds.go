package game

import (
	"fmt"
	"strconv"

	"github.com/DMXMax/mge/chart"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// oddsCmd represents the command to set the odds for the game
var oddsCmd = &cobra.Command{
	Use:     "odds [value]",
	Aliases: []string{"o"},
	Short:   "Set or show the default odds for the game",
	Long:    `Set or show the default odds for the game. This can be a number (0-8) or a name (e.g., "likely", "50/50").`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g == nil {
			return fmt.Errorf("no game selected")
		}
		// if there is an argument, set the odds value
		if len(args) == 1 {
			oddsStr := args[0]
			var newOdds int8

			// try to convert odds to a number. If not, try to match it to a string
			parsed, err := strconv.ParseInt(oddsStr, 10, 8)
			if err != nil { // not a number, try to match it to a string
				matches := chart.MatchOddsPrefix(oddsStr)
				if len(matches) == 0 {
					return fmt.Errorf("invalid odds: %s", oddsStr)
				}
				if len(matches) > 1 { // multiple possible odds
					fmt.Println("Did you mean one of these odds?")
					for _, match := range matches {
						fmt.Printf("%d : %s\n", match, chart.OddsStrList[match])
					}
					return fmt.Errorf("multiple possible odds")
				}
				newOdds = int8(matches[0])
			} else {
				if parsed < 0 || parsed > 8 {
					return fmt.Errorf("odds must be between 0 and 8")
				}
				newOdds = int8(parsed)
			}

			g.SetOdds(newOdds)
			fmt.Printf("Default odds set to %s\n", chart.OddsStrList[newOdds])
			// Persist the change to the database
			if err := db.GamesDB.Save(g).Error; err != nil {
				return fmt.Errorf("failed to save game after changing odds: %w", err)
			}
			return nil
		}

		fmt.Printf("Current default odds: %s\n", chart.OddsStrList[g.Odds])
		return nil
	},
}
