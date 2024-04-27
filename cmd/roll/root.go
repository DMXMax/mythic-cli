package roll

import (
	"fmt"

	"github.com/DMXMax/cli-test/util/dice"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RollCmd = &cobra.Command{
	Use:   "roll",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Die Roll!")
		fmt.Printf("%s\n", dice.RollFate())

		return nil
	},
}
