package game

import (
	"fmt"

	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

var chaos int8

// rootCmd represents the base command when called without any subcommands
var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "create or select a game",
	Long:    `Create a new game with a supplied name. If the game already exists, it will be selected.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("new requires a name")
		}
		name := args[0]
		chaos = 4
		cmd.ParseFlags(args)
		gdb.SetGame(name, chaos)

		return nil
	},
}

func init() {
	newCmd.Flags().Int8Var(&chaos, "chaos", 4, "set the chaos factor for the game")
}
