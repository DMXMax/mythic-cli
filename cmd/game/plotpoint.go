package game

import (
	"math/rand"

	"github.com/DMXMax/mge/util/plot"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// plotPointCmd represents the base command for managing plot points.
var plotPointCmd = &cobra.Command{
	Use:     "plotpoint",
	Aliases: []string{"pp", "plot"},
	Short:   "Manage plot points for the current game",
	Long:    `Manage plot points for the current game. You can add, list, and resolve plot points.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		roll := rand.Intn(100) + 1
		pickTheme := gdb.Current.StoryThemes.GetRandomTheme() // By default, if no subcommand is given, show help.
		pp, err := plot.Chart.GetChartEntry(roll, pickTheme)
		if err != nil {
			return err
		}

		// The main output is just the description
		cmd.Println(pp.Description)

		// All other information is shown with the verbose flag
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			cmd.Printf("\n--- Debug Info ---\n")
			cmd.Printf("Theme: %s\tRoll: %d\n", pickTheme, roll)
			cmd.Printf("%-10s %-10s %-10s %-10s %-10s\n", "Action", "Mystery", "Personal", "Social", "Tension")
			cmd.Printf("%-10d %-10d %-10d %-10d %-10d\n", pp.Action, pp.Mystery, pp.Personal, pp.Social, pp.Tension)
		}
		return nil
	},
}

func init() {
	plotPointCmd.Flags().BoolP("verbose", "v", false, "Show verbose output including theme, roll, and modifiers")
}
