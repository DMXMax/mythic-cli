package game

import (
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var GameCmd = &cobra.Command{
	Use:     "game",
	Aliases: []string{"g"},
	Short:   "Manages Games",
	Long: `Can create and save games in .gob files.
	Games currently have a name and a chaos factor.
	Games have a log or a story that can be added to.
	It will describe the current game if one is selected.
	If one is not selected, it will show the usage.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		g := gdb.Current
		if g != nil {
			cmd.Println("Current Game:", *g)
		} else {
			cmd.Usage()
		}

		return nil
	},
}

func init() {
    GameCmd.AddCommand(createCmd)
    GameCmd.AddCommand(saveCmd)
    GameCmd.AddCommand(chaosCmd)
    GameCmd.AddCommand(loadCmd)
    GameCmd.AddCommand(gameListCmd)
    GameCmd.AddCommand(oddsCmd)
    GameCmd.AddCommand(removeCmd)
}
