package scene

import (
	"fmt"
	"strings"

	"github.com/DMXMax/mge/storage"
	"github.com/DMXMax/mge/util"
	"github.com/DMXMax/mge/util/scene"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// startCmd starts a new scene with an Expected Scene concept.
// It automatically rolls the Chaos Die to determine if the scene is Expected, Altered, or Interrupted.
var startCmd = &cobra.Command{
	Use:   "start <description>",
	Short: "Start a new scene",
	Long: `Start a new scene with an Expected Scene concept. The Chaos Die is automatically rolled
to determine if the scene proceeds as expected, is altered, or is interrupted.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g == nil {
			return fmt.Errorf("no game selected. Use 'game load <name>' to select one")
		}

		// Get Expected Scene concept from arguments
		concept := strings.Join(args, " ")

		// Roll Chaos Die
		rollResult := scene.RollChaosDie(int(g.Chaos))

		// Create scene record
		newScene := storage.Scene{
			GameID:         g.ID,
			Type:           rollResult.SceneType,
			ExpectedConcept: concept,
			ChaosDieRoll:   rollResult.Roll,
			IsActive:       true,
		}

		// Deactivate any existing active scene
		if err := db.GamesDB.Model(&storage.Scene{}).
			Where("game_id = ? AND is_active = ?", g.ID, true).
			Update("is_active", false).Error; err != nil {
			return fmt.Errorf("failed to deactivate existing scene: %w", err)
		}

		// Create new scene
		if err := db.GamesDB.Create(&newScene).Error; err != nil {
			return fmt.Errorf("failed to create scene: %w", err)
		}

		// Display scene type and roll result
		cmd.Printf("Scene Started: %s\n", rollResult.Description)
		cmd.Printf("Expected Scene: %s\n", concept)

		// If Altered or Interrupted, generate Random Event
		if rollResult.SceneType == "altered" || rollResult.SceneType == "interrupt" {
			event := util.GetEvent()
			cmd.Printf("\nRandom Event: %s\n", event.String())

			// Log the event
			eventMsg := fmt.Sprintf("--- Scene Start: %s | Expected: %s | Event: %s ---",
				strings.Title(rollResult.SceneType), concept, event.String())
			entry := gdb.LogEntry{
				Type:  0,
				Msg:   eventMsg,
				GameID: g.ID,
			}
			if err := db.GamesDB.Create(&entry).Error; err != nil {
				return fmt.Errorf("failed to log event: %w", err)
			}
		} else {
			// Log expected scene start
			eventMsg := fmt.Sprintf("--- Scene Start: Expected | %s ---", concept)
			entry := gdb.LogEntry{
				Type:  0,
				Msg:   eventMsg,
				GameID: g.ID,
			}
			if err := db.GamesDB.Create(&entry).Error; err != nil {
				return fmt.Errorf("failed to log scene start: %w", err)
			}
		}

		return nil
	},
}

func init() {
	SceneCmd.AddCommand(startCmd)
}

