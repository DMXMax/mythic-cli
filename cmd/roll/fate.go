package roll

import (
	"fmt"
	"strings"

	"github.com/DMXMax/mythic-cli/util/db"
	"github.com/DMXMax/mythic-cli/util/dice"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

var skillValue int
var difficultyValue int
var opposedRoll bool

// RollFateCmd represents the command to roll Fate dice.
var RollFateCmd = &cobra.Command{
	Use:     "rollfate [message]",
	Aliases: []string{"rf"},
	Short:   "Rolls 4 Fate/Fudge dice (4dF)",
	Long: `Rolls 4 Fate/Fudge dice, which result in a value from -4 to +4.
An optional message can be provided to be included in the log.
An optional skill value can be added to the roll total using the --skill flag.
An optional difficulty can be provided with the --difficulty flag to compare against the total.
Use the --opposed flag to make the difficulty an opposed roll (adds 4dF to the difficulty value).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Roll the dice
		fateRoll := dice.RollFate()

		// Get skill value from flag
		skill, err := cmd.Flags().GetInt("skill")
		if err != nil {
			// This should not happen with properly defined flags
			return fmt.Errorf("failed to get skill flag: %w", err)
		}

		// Get difficulty value from flag
		difficulty, err := cmd.Flags().GetInt("difficulty")
		if err != nil {
			return fmt.Errorf("failed to get difficulty flag: %w", err)
		}

		// Get opposed flag value
		isOpposed, err := cmd.Flags().GetBool("opposed")
		if err != nil {
			return fmt.Errorf("failed to get opposed flag: %w", err)
		}

		// Add skill as a modifier if it's not the default value
		if cmd.Flags().Changed("skill") {
			fateRoll.Modifiers = append(fateRoll.Modifiers, dice.RollModifier{
				Mod:         int8(skill),
				Description: "skill",
			})
		}

		rollTotal := fateRoll.Total()
		message := strings.Join(args, " ")
		var logMessage string

		rollStr := fateRoll.String()
		if cmd.Flags().Changed("skill") {
			rollStr = fmt.Sprintf("%s; skill %d -> %d", rollStr, skill, rollTotal)
		}

		if cmd.Flags().Changed("difficulty") {
			finalDifficulty := difficulty
			var outcome int
			if isOpposed {
				opposedRoll := dice.RollFate()
				finalDifficulty = difficulty + opposedRoll.Total()
				outcome = rollTotal - finalDifficulty
				opponentStr := fmt.Sprintf("%s; skill %d -> %d", opposedRoll.String(), difficulty, finalDifficulty)
				rollStr = fmt.Sprintf("%s vs Opponent (%s): %s (%+d)", rollStr, opponentStr, getOutcomeString(outcome), outcome)
			} else {
				outcome = rollTotal - difficulty
				rollStr = fmt.Sprintf("%s vs diff %d: %s (%+d)", rollStr, difficulty, getOutcomeString(outcome), outcome)
			}
		}

		if strings.TrimSpace(message) != "" {
			logMessage = fmt.Sprintf("%s | 4dF %s", message, rollStr)
		} else {
			logMessage = fmt.Sprintf("4dF %s", rollStr)
		}

		fmt.Println(logMessage)

		if gdb.Current != nil {
			entry := gdb.LogEntry{Type: 0, Msg: logMessage, GameID: gdb.Current.ID}
			if err := db.GamesDB.Create(&entry).Error; err != nil {
				return fmt.Errorf("failed to save game after fate roll: %w", err)
			}
		}

		return nil
	},
}

func init() {
	RollFateCmd.Flags().IntVarP(&skillValue, "skill", "s", 0, "skill value to add to the roll")
	RollFateCmd.Flags().IntVarP(&difficultyValue, "difficulty", "d", 0, "difficulty to compare the roll against")
	RollFateCmd.Flags().BoolVarP(&opposedRoll, "opposed", "o", false, "make the difficulty an opposed roll")
}

func getOutcomeString(outcome int) string {
	if outcome >= 3 {
		return "Success With Style"
	} else if outcome > 0 {
		return "Success"
	}
	if outcome < 0 {
		return "Fail"
	}
	return "Tie"
}
