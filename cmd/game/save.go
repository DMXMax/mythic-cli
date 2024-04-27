package game

import (
	"encoding/gob"
	"fmt"
	"os"

	gdb "github.com/DMXMax/cli-test/util/game"
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
		filename := g.Name + ".gob"
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		enc := gob.NewEncoder(file)
		err = enc.Encode(g)
		if err != nil {
			return err
		}
		fmt.Printf("Game saved to %s\n", filename)

		return nil
	},
}

func init() {

}
