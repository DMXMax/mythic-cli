package game

import (
	"fmt"

	gdb "github.com/DMXMax/cli-test/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "create or select a game",
	Long:  `Create a new game with a supplied name. If the game already exists, it will be selected.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("new requires a name")
		}
		name := args[0]

		gdb.SetGame(name)

		return nil
	},
}
