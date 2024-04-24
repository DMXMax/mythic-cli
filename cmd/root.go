/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cli-test/cmd/env"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-test",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Hello, World!")
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
		for {
			cmd.Print("shell> ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanLines)
			scanner.Scan()
			cmd.Print([]string{scanner.Text()})
			cmd.Print("You entered: ", []string{scanner.Text()}, "\n")
			if scanner.Err() != nil {
				return scanner.Err()
			}
			text := strings.Fields(scanner.Text())
			newCmd, args, err := cmd.Find(text)
			//fmt.Println("newCmd: ", newCmd, "args: ", args, "err: ", err)
			if err != nil {
				cmd.Println(err)
				continue
			}
			if newCmd == cmd {
				continue
			}
			fmt.Println("newCmd: ", newCmd.Name(), "args: ", args)
			//newCmd.SetArgs(args)
			err = newCmd.RunE(newCmd, args)

			if err != nil {
				cmd.Println(err)
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
	Long: `Quit the shell, but longer",
`,
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
	shellCmd.AddCommand(shellQuitCmd)
	env.EnvCmd.AddCommand(env.AddCmd)
	shellCmd.AddCommand(env.EnvCmd)
	rootCmd.AddCommand(shellCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
