// Package scene provides commands for managing scenes in Mythic games.
package scene

import (
	"github.com/spf13/cobra"
)

// SceneCmd is the root command for scene management.
var SceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "Manage game scenes",
	Long: `Manage scenes in your Mythic game. Scenes represent distinct moments or locations in your game narrative.

When you start a scene, the Chaos Die is automatically rolled to determine if the scene proceeds as expected,
is altered, or is interrupted. Altered and Interrupted scenes generate Random Events.`,
}
