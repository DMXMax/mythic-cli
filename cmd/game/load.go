package game

import (
	"fmt"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load a game from a file",
	Long:  `Load a game from a given file. `,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Cobra handles flag parsing automatically
		if gameName == "" {
			return fmt.Errorf("no file specified")
		}

		g := &gdb.Game{Name: gameName}
		result := db.GamesDB.Preload("Log").Where(g).First(g)

		if result.Error == nil {
			gdb.Current = g
			cmd.Printf("Loaded game: %s (Chaos: %d)\n", g.Name, g.Chaos)
		} else {
			return fmt.Errorf("could not load game '%s': %w", gameName, result.Error)
		}

		return nil
	},
}
var gameName string

var gameListCmd = &cobra.Command{
	Use:   "list",
	Short: "list current games",
	Long:  `List the games in the database `,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Cobra handles flag parsing automatically
		games := []gdb.Game{}
		result := db.GamesDB.Find(&games)

		if result.Error != nil {
			return result.Error
		}

		if len(games) > 0 {
			fmt.Println("Games Available:")
		} else {
			fmt.Println("No Games Available")
		}

		for _, game := range games {
			fmt.Printf("\t %s\n", game.Name)
		}

		return nil
	},
}

func init() {
	loadCmd.Flags().StringVar(&gameName, "name", "", "load game from this game name")
}
