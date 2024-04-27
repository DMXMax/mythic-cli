package game

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var GameCmd = &cobra.Command{
	Use:   "game",
	Short: "commands about games",
	Long:  `Create New, Save, and Load Games`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	GameCmd.AddCommand(NewCmd)
	GameCmd.AddCommand(SaveCmd)
}
