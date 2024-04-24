package env

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var EnvCmd = &cobra.Command{
	Use:   "env",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello, World!")

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
