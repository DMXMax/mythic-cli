package log

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/DMXMax/mythic-cli/util/db"
    gdb "github.com/DMXMax/mythic-cli/util/game"
    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var LogCmd = &cobra.Command{
    Use:     "gamelog",
    Aliases: []string{"s", "gl", "log"},
    Short:   "manage game logs",
    Long:    `Create New, Save, and Load logs`,
    RunE: func(cmd *cobra.Command, args []string) error {
        // Default behavior: print logs, optionally limited by a number
        return runPrint(args)
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
		if err := db.GamesDB.Save(g).Error; err != nil {
			return fmt.Errorf("failed to save game after adding log: %w", err)
		}
		fmt.Println("Log entry added and game saved.")

		return nil
	},
}

var printCmd = &cobra.Command{
    Use:     "print [n]",
    Aliases: []string{"p"},
    Short:   "print out story log",
    Long:    `Print out the story log. Optionally provide a number to print that many recent entries (most recent shown last).`,
    RunE: func(cmd *cobra.Command, args []string) error {
        return runPrint(args)
    },
}

var removeLogCmd = &cobra.Command{
	Use:     "remove [n]",
	Aliases: []string{"rm"},
	Short:   "remove last n log entries",
	Long:    "Remove the last n log entries. If n is not provided, removes the last one.",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		logLen := len(g.Log)
		if n > logLen {
			fmt.Printf("Cannot remove %d entries, only %d exist. Removing all %d entries.\n", n, logLen, logLen)
			n = logLen
		}

		if n == 0 {
			fmt.Println("No log entries to remove.")
			return nil
		}

		entriesToRemove := g.Log[logLen-n:]

		if err := db.GamesDB.Delete(&entriesToRemove).Error; err != nil {
			return fmt.Errorf("failed to remove log entries from database: %w", err)
		}

		g.Log = g.Log[:logLen-n]
		fmt.Printf("Removed last %d log entry(s).\n", n)
		return nil
	},
}

func init() {
    LogCmd.AddCommand(AddGameLogCmd)
    LogCmd.AddCommand(printCmd)
    LogCmd.AddCommand(removeLogCmd)
}

// runPrint implements the actual printing logic shared by `log` and `log print`.
// If args[0] is a positive integer, prints that many most recent entries; otherwise prints a default number.
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
