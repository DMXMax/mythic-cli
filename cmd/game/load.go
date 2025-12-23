package game

import (
	"fmt"
	"strings"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// loadCmd loads a game by name and sets it as the current game.
// The game name can be provided as a positional argument or via the --name flag.
var loadCmd = &cobra.Command{
	Use:   "load [name]",
	Short: "Load a game by name",
	Long:  `Load a game by name and set it as the current game. You can pass the name as a positional argument or via --name.`,
	RunE: func(cmd *cobra.Command, args []string) error {
        // Accept either positional name or --name flag for convenience
        // Join all args to handle multi-word names (e.g., "Kat in Shadow")
        var name string
        if len(args) > 0 {
            name = strings.Join(args, " ")
        } else {
            name = gameName
        }
        name = strings.TrimSpace(name)
        if name == "" {
            return fmt.Errorf("no game name specified")
        }

        g := &gdb.Game{Name: name}
        result := db.GamesDB.Where(g).First(g)

        if result.Error == nil {
            gdb.Current = g
            cmd.Printf("Loaded game: %s (Chaos: %d)\n", g.Name, g.Chaos)
        } else {
            return fmt.Errorf("could not load game '%s': %w", name, result.Error)
        }

        return nil
    },
}
var gameName string

// gameListCmd lists all games stored in the database.
var gameListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all games in the database",
	Long:  `List all games stored in the database. Shows the names of all available games.`,
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
