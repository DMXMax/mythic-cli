// Package descriptor provides commands for generating descriptors from Elements Meaning Tables.
package descriptor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// DescriptorCmd is the root command for descriptor generation.
var DescriptorCmd = &cobra.Command{
	Use:   "descriptor [type] [number]",
	Short: "Generate descriptors from Elements Meaning Tables",
	Long: `Generate descriptors from Elements Meaning Tables. These tables provide words and phrases
for describing characters, locations, objects, and other game elements.

Use 'descriptor list' to see all available descriptor tables.
Use 'descriptor <type> [number]' to generate descriptors from a specific table.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If no args, show help
		if len(args) == 0 {
			return cmd.Help()
		}
		
		// Try to match table name by progressively combining arguments
		var tableType string
		var countArgIndex int
		found := false
		
		// Try matching progressively longer argument sequences
		for i := 1; i <= len(args); i++ {
			candidate := strings.ToLower(strings.TrimSpace(strings.Join(args[0:i], " ")))
			if isTableAvailable(candidate) {
				tableType = candidate
				countArgIndex = i
				found = true
				break
			}
		}
		
		if !found {
			// No match found, show available tables
			cmd.Printf("Unknown table type: %s\n\n", strings.Join(args, " "))
			cmd.Println("Available tables:")
			tables := getAvailableTables()
			for _, table := range tables {
				cmd.Printf("  %s\n", table)
			}
			return fmt.Errorf("use 'descriptor list' to see all available tables")
		}
		
		// Get number from remaining arguments (default: 1)
		count := 1
		if countArgIndex < len(args) {
			var err error
			count, err = strconv.Atoi(args[countArgIndex])
			if err != nil || count < 1 {
				return fmt.Errorf("number must be a positive integer, got: %s", args[countArgIndex])
			}
			if count > 20 {
				return fmt.Errorf("number cannot exceed 20")
			}
		}
		
		// Generate entries
		entries := getRandomEntries(tableType, count)
		
		if len(entries) == 0 {
			return fmt.Errorf("failed to generate entries from table '%s'", tableType)
		}
		
		// Display results
		for i, entry := range entries {
			cmd.Println(entry)
			if i < len(entries)-1 {
				// Add spacing between entries if generating multiple
				if count > 1 {
					cmd.Println()
				}
			}
		}
		
		return nil
	},
}

