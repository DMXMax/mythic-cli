package log

import (
	"fmt"
	"strings"

	gdb "github.com/DMXMax/cli-test/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var LogCmd = &cobra.Command{
	Use:     "gamelog",
	Aliases: []string{"s", "gl", "log"},
	Short:   "manage game logs",
	Long:    `Create New, Save, and Load logs`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

var AddGameLogCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "add to game log",
	Long:    `Add Entry to game story log. `,
	RunE: func(cmd *cobra.Command, args []string) error {

		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		g.AddtoGameLog(0, strings.Join(args, " "))
		return nil

	},
}

var dumpCmd = &cobra.Command{
	Use:     "print",
	Aliases: []string{"p"},
	Short:   "print out story log",
	Long:    `Print out the story log, with an optional depth`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		for _, s := range g.GameLog {
			fmt.Println(s)
		}
		return nil

	},
}

var detailsCmd = &cobra.Command{
	Use:     "details",
	Aliases: []string{"d"},
	Short:   "print out story settings",
	Long:    `Print out the story settings`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		fmt.Println(g.GameLog)
		return nil

	},
}

func init() {
	LogCmd.AddCommand(AddGameLogCmd)
	LogCmd.AddCommand(dumpCmd)
	LogCmd.AddCommand(detailsCmd)
	//StoryCmd.AddCommand(newCmd)

}
