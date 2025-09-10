/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/DMXMax/mge/chart"

	"github.com/DMXMax/mythic-cli/cmd/scene"
	gdb "github.com/DMXMax/mythic-cli/util/game"

	"github.com/DMXMax/mythic-cli/cmd/game"
	gamelog "github.com/DMXMax/mythic-cli/cmd/log"
	"github.com/DMXMax/mythic-cli/cmd/roll"

	"github.com/DMXMax/mythic-cli/util/input"
	"github.com/peterh/liner"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mythic-cli",
	Short: "Mythic Game Master Emulator CLI",
	Long: `A command-line tool for solo RPG gaming using the Mythic Game Master Emulator system.

Features:
- Interactive shell for game management
- Dice rolling with Mythic fate chart
- Game state persistence with SQLite
- Story logging and chaos factor management
- Character and scene management

Perfect for solo RPG adventures, GM-less gaming, and story generation.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Usage()
		return nil
	},
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A simple shell, holding token to use interactive commands",
	Long: `A simple shell, holding token to use interactive commands -- the long version
    Examples go here.
`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		// Use liner to get arrow-key history and line editing
		l := liner.NewLiner()
		defer l.Close()
		l.SetCtrlCAborts(true)
		// Make liner available to commands for sub-prompts
		input.SetPrompter(l)

		// Load/save persistent history
		home, _ := os.UserHomeDir()
		histPath := filepath.Join(home, ".mythic-cli_history")
		if f, err := os.Open(histPath); err == nil {
			_, _ = l.ReadHistory(f)
			_ = f.Close()
		}
		defer func() {
			if f, err := os.Create(histPath); err == nil {
				_, _ = l.WriteHistory(f)
				_ = f.Close()
			}
		}()

		for {
			var prompt string
			if gdb.Current != nil {
				g := gdb.Current
				oddsName := chart.OddsStrList[g.Odds]
				prompt = fmt.Sprintf("%s (C:%d O:%s)> ", g.Name, g.Chaos, oddsName)
			} else {
				prompt = "shell> "
			}

			line, err := l.Prompt(prompt)
			if err == liner.ErrPromptAborted { // Ctrl-C
				cmd.Println("Goodbye!")
				return nil
			}
			if err == io.EOF { // Ctrl-D
				cmd.Println("Goodbye!")
				return nil
			}
			if err != nil {
				return err
			}

			input := strings.TrimSpace(line)
			if input == "" {
				continue
			}

			fields := strings.Fields(input)
			newCmd, newArgs, err := cmd.Find(fields)
			if err != nil {
				cmd.Println(err)
				continue
			}
			// If Find returns the same command, it means no subcommand was found.
			if newCmd == cmd {
				cmd.Printf("Error: unknown command \"%s\" for \"%s\"\n", fields[0], cmd.CommandPath())
				continue
			}

			// Check if help is requested
			hasHelp := false
			for _, arg := range newArgs {
				if arg == "--help" || arg == "-h" {
					hasHelp = true
					break
				}
			}

			if hasHelp {
				// Show help for the command
				newCmd.Help()
			} else {
				// Set the args for the command and execute it normally
				newCmd.SetArgs(newArgs)

				// Parse flags to ensure default values are set
				if err := newCmd.Flags().Parse(newArgs); err != nil {
					cmd.Println(err)
					continue
				}

				if newCmd.RunE != nil {
					// After parsing, the non-flag arguments are available via Flags().Args()
					if err := newCmd.RunE(newCmd, newCmd.Flags().Args()); err != nil {
						cmd.Println(err)
					}
				} else if newCmd.Run != nil {
					newCmd.Run(newCmd, newCmd.Flags().Args())
					l.AppendHistory(input)
				}

				// Reset flags on the executed command to avoid carry-over in the shell
				newCmd.Flags().VisitAll(func(f *pflag.Flag) {
					f.Value.Set(f.DefValue)
					f.Changed = false
				})
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var shellQuitCmd = &cobra.Command{
	Use:   "quit",
	Short: "Quit the shell",
	Long:  `Quit the shell, but longer`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Goodbye!")
		os.Exit(0)
		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-test.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	shellCmd.AddCommand(shellQuitCmd, scene.SceneCmd, game.GameCmd,
		roll.RollCmd, gamelog.LogCmd, shellHelpCommand)

	rootCmd.AddCommand(shellCmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//shellCmd.SetUsageTemplate(Template)
}

var shellHelpCommand = &cobra.Command{
	Use:   "help",
	Short: "Shell help",
	Long:  `Provides detailed shell help help`,
	RunE: func(cmd *cobra.Command, args []string) error {
		newcmd, _, err := cmd.Parent().Find(args)
		if err != nil {
			return err
		}
		if newcmd == cmd {
			return cmd.Parent().Usage()
		}
		return newcmd.Help()
	},
}
