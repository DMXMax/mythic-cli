package game

import (
	"fmt"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save a game to a file",
	Long:  `Save a game to a given file name. Currently it will overwrite any existing file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		/*		if g.ID == 0 {
				db.GamesDB.Create(g)
			} else {*/
		db.GamesDB.Save(g)
		//}
		//fmt.Printf("Game saved as %d", g.ID)

		return nil
	},
}

func init() {

}
