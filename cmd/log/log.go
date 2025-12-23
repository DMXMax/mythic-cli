// Package log provides commands for managing game story logs.
// Logs contain all dice rolls, story events, and narrative entries for a game.
package log

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DMXMax/mythic-cli/util/db"
	gdb "github.com/DMXMax/mythic-cli/util/game"
	"github.com/spf13/cobra"
)

// LogCmd is the root command for managing game logs.
// When invoked without subcommands, it prints recent log entries (default: last 20).
// An optional number can be provided to limit the number of entries shown.
var LogCmd = &cobra.Command{
	Use:     "gamelog",
	Aliases: []string{"s", "gl", "log"},
	Short:   "Manage game logs",
	Long:    `Manage game story logs including viewing, adding, and removing entries.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If an argument is provided, validate it's a number before calling runPrint
		if len(args) > 0 {
			// Check if it's a valid number
			if _, err := strconv.Atoi(args[0]); err != nil {
				// Not a number - suggest valid subcommands
				return fmt.Errorf("unknown argument: %q\n\nAvailable subcommands:\n  add, print, remove\n\nUse \"log <number>\" to print that many entries, or \"log --help\" for more information", args[0])
			}
		}
		// Default behavior: print logs, optionally limited by a number
		return runPrint(args)
	},
}

// AddGameLogCmd adds a new entry to the current game's story log.
// The entry is immediately persisted to the database.
var AddGameLogCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add an entry to the game log",
	Long:    `Add an entry to the game story log. The entry is immediately saved to the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		msg := strings.Join(args, " ")
		// Create log entry directly in database to avoid duplicates
		entry := gdb.LogEntry{Type: 0, Msg: msg, GameID: g.ID}
		if err := db.GamesDB.Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to save log entry: %w", err)
		}
		fmt.Println("Log entry added and game saved.")

		return nil
	},
}

// printCmd prints recent log entries from the current game.
// An optional number can be provided to limit the number of entries (default: 20).
// Entries are displayed in chronological order (oldest first).
var printCmd = &cobra.Command{
	Use:     "print [n]",
	Aliases: []string{"p", "list", "l"},
	Short:   "Print recent log entries",
	Long:    `Print out the story log. Optionally provide a number to print that many recent entries (most recent shown last).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if help was requested
		if len(args) > 0 && (args[0] == "help" || args[0] == "--help" || args[0] == "-h") {
			return cmd.Help()
		}
		return runPrint(args)
	},
}

// removeLogCmd removes the last n log entries from the current game.
// If no number is provided, it removes the last entry.
// The entries are permanently deleted from the database.
var removeLogCmd = &cobra.Command{
	Use:     "remove [n]",
	Aliases: []string{"rm"},
	Short:   "Remove the last n log entries",
	Long:    "Remove the last n log entries. If n is not provided, removes the last one.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if help was requested
		if len(args) > 0 && (args[0] == "help" || args[0] == "--help" || args[0] == "-h") {
			return cmd.Help()
		}

		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current

		n := 1
		var err error
		if len(args) > 0 {
			n, err = strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid number: %w", err)
			}
		}

		if n <= 0 {
			return fmt.Errorf("number of entries to remove must be positive")
		}

		// Fetch the last n entries to be removed
		var entriesToRemove []gdb.LogEntry
		q := db.GamesDB.Model(&gdb.LogEntry{}).
			Where("game_id = ?", g.ID).
			Order("created_at DESC").
			Limit(n)
		if err := q.Find(&entriesToRemove).Error; err != nil {
			return fmt.Errorf("failed to load log entries for removal: %w", err)
		}

		numToRemove := len(entriesToRemove)
		if numToRemove == 0 {
			fmt.Println("No log entries to remove.")
			return nil
		}

		if numToRemove < n {
			fmt.Printf("Cannot remove %d entries, only %d exist. Removing all %d entries.\n", n, numToRemove, numToRemove)
		}

		if err := db.GamesDB.Delete(&entriesToRemove).Error; err != nil {
			return fmt.Errorf("failed to remove log entries from database: %w", err)
		}

		// Invalidate the in-memory log to force a reload on next `log print`
		g.Log = nil
		fmt.Printf("Removed last %d log entry(s).\n", numToRemove)
		return nil
	},
}

func init() {
	LogCmd.AddCommand(AddGameLogCmd)
	LogCmd.AddCommand(printCmd)
	LogCmd.AddCommand(removeLogCmd)
}

// runPrint implements the actual printing logic shared by `log` and `log print`.
// It fetches the most recent n entries from the database and displays them in chronological order.
// If args[0] is a positive integer, it prints that many most recent entries; otherwise prints a default number (20).
func runPrint(args []string) error {
	if gdb.Current == nil {
		return fmt.Errorf("no game selected")
	}
	g := gdb.Current

	// Default to a recent window if no number is specified
	n := 20
	var err error
	if len(args) > 0 {
		n, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid number: %w", err)
		}
	}
	if n <= 0 {
		return fmt.Errorf("number of entries to print must be positive")
	}

	// Fetch the last n entries from DB ordered by newest first
	var entries []gdb.LogEntry
	q := db.GamesDB.Model(&gdb.LogEntry{}).
		Where("game_id = ?", g.ID).
		Order("created_at DESC").
		Limit(n)
	if err := q.Find(&entries).Error; err != nil {
		return fmt.Errorf("failed to load log entries: %w", err)
	}

	// Print oldest-first for natural reading by reversing the slice
	for i := len(entries) - 1; i >= 0; i-- {
		s := entries[i]
		fmt.Printf("%s - %s\n", s.CreatedAt.Format("2006-01-02 15:04:05"), s.Msg)
	}

	return nil
}
