package scene

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var SceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Scene Variable!")

		return nil
	},
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Add Variable!")

		return nil
	},
}
