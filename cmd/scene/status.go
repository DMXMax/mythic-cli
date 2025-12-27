package scene

import (
	"fmt"
	"strings"

	"github.com/DMXMax/mge/chart"
	"github.com/DMXMax/mge/storage"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// statusCmd displays the current scene status.
var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "Show current scene status",
	Long:    `Displays information about the current active scene, including type, concept, and roll result.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g == nil {
			return fmt.Errorf("no game selected. Use 'game load <name>' to select one")
		}

		// Find active scene
		var currentScene storage.Scene
		result := db.GamesDB.Where("game_id = ? AND is_active = ?", g.ID, true).
			First(&currentScene)

		if result.Error != nil {
			cmd.Println("No active scene.")
			return nil
		}

		// Display scene information
		cmd.Printf("Current Scene:\n")
		cmd.Printf("  Type: %s\n", strings.Title(currentScene.Type))
		cmd.Printf("  Expected Concept: %s\n", currentScene.ExpectedConcept)
		cmd.Printf("  Chaos Die Roll: %d (Chaos: %d)\n", currentScene.ChaosDieRoll, chart.ChaosInternalToUser(int(g.Chaos)))
		return nil
	},
}

func init() {
	SceneCmd.AddCommand(statusCmd)
}
