package game

import (
	"fmt"

	"github.com/DMXMax/cli-test/util/db"
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
		if gameName == "" {
			return fmt.Errorf("no file specified")
		}

		/*fileName := fmt.Sprintf("%s.gob", fileName)
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
		g =
		fmt.Printf("Loaded from %s\n", fileName)*/
		g := &gdb.Game{Name: gameName}
		result := db.GamesDB.Where(g).Preload("GameLog").First(g)

		if result.Error == nil {
			gdb.Current = g
		} else {
			fmt.Println(result.Error)
		}

		return nil
	},
}
var gameName string

func init() {
	loadCmd.Flags().StringVar(&gameName, "name", "", "load game from this game name")
}
