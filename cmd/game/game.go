package game

import (
	gdb "github.com/DMXMax/cli-test/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var GameCmd = &cobra.Command{
	Use:     "game",
	Aliases: []string{"g", "story"},
	Short:   "commands about games",
	Long:    `Create New, Save, and Load Games`,
	RunE: func(cmd *cobra.Command, args []string) error {

		g := gdb.Current
		if g != nil {
			cmd.Println("Current Game:", *g)
		} else {
			cmd.Println("No game selected")
		}

		return nil
	},
}

func init() {
	GameCmd.AddCommand(newCmd)
	GameCmd.AddCommand(saveCmd)
	GameCmd.AddCommand(chaosCmd)
}
