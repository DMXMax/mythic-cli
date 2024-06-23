package game

import (
	"encoding/gob"
	"fmt"
	"os"

	gdb "github.com/DMXMax/cli-test/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load a game from a file",
	Long:  `Load a game from a given file. `,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.ParseFlags(args)
		if err != nil {
			return err
		}
		if fileName == "" {
			return fmt.Errorf("no file specified")
		}

		fileName := fmt.Sprintf("%s.gob", fileName)
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		dec := gob.NewDecoder(file)
		g := gdb.Game{}
		err = dec.Decode(&g)
		if err != nil {
			return err
		}
		gdb.Current = &g

		fmt.Printf("Loaded from %s\n", fileName)

		return nil
	},
}
var fileName string

func init() {
	loadCmd.Flags().StringVar(&fileName, "file", "", "load game from this file")
}
