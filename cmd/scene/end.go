package scene

import (
	"fmt"

	"github.com/DMXMax/mge/storage"
	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// endCmd ends the current active scene.
// This deactivates the scene but does not handle Threads/Characters Lists updates
// (those should be done via the scene_end command/tool separately).
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End the current scene",
	Long: `Ends the current active scene. This deactivates the scene in the database.
Note: To update Threads/Characters Lists and adjust Chaos Factor, use the scene_end command separately.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g := gdb.Current
		if g == nil {
			return fmt.Errorf("no game selected. Use 'game load <name>' to select one")
		}

		// Find and deactivate active scene
		result := db.GamesDB.Model(&storage.Scene{}).
			Where("game_id = ? AND is_active = ?", g.ID, true).
			Update("is_active", false)

		if result.Error != nil {
			return fmt.Errorf("failed to end scene: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			cmd.Println("No active scene to end.")
			return nil
		}

		cmd.Println("Scene ended.")

		// Log scene end
		entry := gdb.LogEntry{
			Type:  0,
			Msg:   "--- Scene End ---",
			GameID: g.ID,
		}
		if err := db.GamesDB.Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to log scene end: %w", err)
		}

		return nil
	},
}

func init() {
	SceneCmd.AddCommand(endCmd)
}

