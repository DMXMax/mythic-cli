// Package scene provides commands for managing scenes in Mythic games.
// Scenes represent distinct moments or locations in your game narrative.
package scene

import (
	"fmt"
	"strings"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// SceneCmd is the root command for scene management.
var SceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "Manage game scenes",
	Long: `Manage scenes in your Mythic game.

Scenes represent distinct moments or locations in your game narrative.
Use "scene start" to begin a new scene and "scene end" to mark the end of the current scene.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// startCmd creates a scene start log entry.
var startCmd = &cobra.Command{
	Use:   "start [name]",
	Short: "Start a new scene",
	Long:  `Start a new scene by creating a scene start marker in the game log. The scene name is optional but recommended.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}

		sceneName := strings.TrimSpace(strings.Join(args, " "))
		if sceneName == "" {
			sceneName = "Untitled Scene"
		}

		entry := gdb.LogEntry{
			Type:   gdb.LogTypeSceneStart,
			Msg:    sceneName,
			GameID: gdb.Current.ID,
		}

		if err := db.GamesDB.Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to create scene start entry: %w", err)
		}

		fmt.Printf("Scene started: %s\n", sceneName)
		return nil
	},
}

// endCmd creates a scene end log entry.
var endCmd = &cobra.Command{
	Use:   "end [notes]",
	Short: "End the current scene",
	Long:  `End the current scene by creating a scene end marker in the game log. Optional notes can be provided to summarize the scene.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}

		notes := strings.TrimSpace(strings.Join(args, " "))
		if notes == "" {
			notes = "Scene ended"
		}

		entry := gdb.LogEntry{
			Type:   gdb.LogTypeSceneEnd,
			Msg:    notes,
			GameID: gdb.Current.ID,
		}

		if err := db.GamesDB.Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to create scene end entry: %w", err)
		}

		fmt.Printf("Scene ended: %s\n", notes)
		return nil
	},
}

func init() {
	SceneCmd.AddCommand(startCmd, endCmd)
}
