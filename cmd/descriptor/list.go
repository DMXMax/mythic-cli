package descriptor

import (
	"sort"
	"strings"

	"github.com/DMXMax/mge/util/elements"
	"github.com/spf13/cobra"
)

// listCmd lists all available descriptor/element tables.
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all available descriptor tables",
	Long:    `Lists all available Elements Meaning Tables that can be used for generating descriptors.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tables := getAvailableTables()
		
		// Sort for consistent output
		sort.Strings(tables)
		
		cmd.Println("Available Descriptor Tables:")
		cmd.Println()
		
		for _, table := range tables {
			cmd.Printf("  %s\n", table)
		}
		
		cmd.Printf("\nTotal: %d tables\n", len(tables))
		cmd.Println("\nUse 'descriptor <type> [number]' to generate entries from a table.")
		
		return nil
	},
}

func init() {
	DescriptorCmd.AddCommand(listCmd)
}

// getAvailableTables returns a list of all available element table names.
func getAvailableTables() []string {
	tables := make([]string, 0)
	
	// Add generic descriptor tables
	tables = append(tables, "descriptors", "descriptors1", "descriptors2")
	
	// Add all element tables that contain "descriptor" in their name or are commonly used
	allTables := map[string]bool{
		"descriptors":                true,
		"descriptors1":               true,
		"descriptors2":               true,
		"characters":                 true,
		"locations":                  true,
		"objects":                    true,
		"actions":                    true,
		"actions2":                   true,
		"adventure tone":             true,
		"alien species descriptors": true,
		"animal actions":            true,
		"army descriptors":          true,
		"cavern descriptors":        true,
		"character actions combat":  true,
		"character actions general": true,
		"character appearance":      true,
		"character background":     true,
		"character conversations":   true,
		"character identity":       true,
		"character motivations":    true,
		"character personality":    true,
		"character skills":         true,
		"character traits flaws":   true,
		"city descriptors":         true,
		"civilization descriptors":  true,
		"creature abilities":       true,
		"creature descriptors":      true,
		"cryptic message":          true,
		"curses":                   true,
		"domicile descriptors":     true,
		"dungeon descriptors":      true,
		"dungeon traps":            true,
		"forest descriptors":       true,
		"gods":                     true,
		"legends":                  true,
		"magic item descriptors":   true,
		"mutation descriptors":     true,
		"names":                    true,
		"noble house":              true,
		"plot twists":              true,
		"powers":                  true,
		"scavenging results":       true,
		"smells":                  true,
		"sounds":                  true,
		"spell effects":            true,
		"starship descriptors":     true,
		"terrain descriptors":      true,
		"undead descriptors":       true,
		"visions dreams":          true,
	}
	
	// Check which tables actually exist by trying to access them
	for tableName := range allTables {
		if isTableAvailable(tableName) {
			// Avoid duplicates
			alreadyAdded := false
			for _, t := range tables {
				if strings.EqualFold(t, tableName) {
					alreadyAdded = true
					break
				}
			}
			if !alreadyAdded {
				tables = append(tables, tableName)
			}
		}
	}
	
	return tables
}

// isTableAvailable checks if an element table is available.
func isTableAvailable(tableName string) bool {
	tableName = strings.ToLower(strings.TrimSpace(tableName))
	
	// Special cases for generic tables
	if tableName == "actions" || tableName == "actions2" {
		return len(elements.ActionTable1) > 0 || len(elements.ActionTable2) > 0
	}
	if tableName == "descriptors" || tableName == "descriptors1" {
		return len(elements.Descriptor1) > 0
	}
	if tableName == "descriptors2" {
		return len(elements.Descriptor2) > 0
	}
	
	// Check registered tables (we'll check a few common ones)
	switch tableName {
	case "characters":
		return len(elements.CharacterDescriptors) > 0
	case "locations":
		return len(elements.LocationTable) > 0
	case "objects":
		return len(elements.ObjectDescriptors) > 0
	case "adventure tone":
		return len(elements.AdventureToneTable) > 0
	case "alien species descriptors":
		return len(elements.AlienSpeciesDescriptorsTable) > 0
	case "animal actions":
		return len(elements.AnimalActionsTable) > 0
	case "army descriptors":
		return len(elements.ArmyDescriptorsTable) > 0
	case "cavern descriptors":
		return len(elements.CavernDescriptorsTable) > 0
	case "character actions combat":
		return len(elements.CharacterActionsCombatTable) > 0
	case "character actions general":
		return len(elements.CharacterActionsGeneralTable) > 0
	case "character appearance":
		return len(elements.CharacterAppearanceTable) > 0
	case "character background":
		return len(elements.CharacterBackgroundTable) > 0
	case "character conversations":
		return len(elements.CharacterConversationsTable) > 0
	case "character identity":
		return len(elements.CharacterIdentityTable) > 0
	case "character motivations":
		return len(elements.CharacterMotivationsTable) > 0
	case "character personality":
		return len(elements.CharacterPersonalityTable) > 0
	case "character skills":
		return len(elements.CharacterSkillsTable) > 0
	case "character traits flaws":
		return len(elements.CharacterTraitsFlawsTable) > 0
	case "city descriptors":
		return len(elements.CityDescriptorsTable) > 0
	case "civilization descriptors":
		return len(elements.CivilizationDescriptorsTable) > 0
	case "creature abilities":
		return len(elements.CreatureAbilitiesTable) > 0
	case "creature descriptors":
		return len(elements.CreatureDescriptorsTable) > 0
	case "cryptic message":
		return len(elements.CrypticMessageTable) > 0
	case "curses":
		return len(elements.CursesTable) > 0
	case "domicile descriptors":
		return len(elements.DomicileDescriptorsTable) > 0
	case "dungeon descriptors":
		return len(elements.DungeonDescriptorsTable) > 0
	case "dungeon traps":
		return len(elements.DungeonTrapsTable) > 0
	case "forest descriptors":
		return len(elements.ForestDescriptorsTable) > 0
	case "gods":
		return len(elements.GodsTable) > 0
	case "legends":
		return len(elements.LegendsTable) > 0
	case "magic item descriptors":
		return len(elements.MagicItemDescriptorsTable) > 0
	case "mutation descriptors":
		return len(elements.MutationDescriptorsTable) > 0
	case "names":
		return len(elements.NamesTable) > 0
	case "noble house":
		return len(elements.NobleHouseTable) > 0
	case "plot twists":
		return len(elements.PlotTwistsTable) > 0
	case "powers":
		return len(elements.PowersTable) > 0
	case "scavenging results":
		return len(elements.ScavengingResultsTable) > 0
	case "smells":
		return len(elements.SmellsTable) > 0
	case "sounds":
		return len(elements.SoundsTable) > 0
	case "spell effects":
		return len(elements.SpellEffectsTable) > 0
	case "starship descriptors":
		return len(elements.StarshipDescriptorsTable) > 0
	case "terrain descriptors":
		return len(elements.TerrainDescriptorsTable) > 0
	case "undead descriptors":
		return len(elements.UndeadDescriptorsTable) > 0
	case "visions dreams":
		return len(elements.VisionsDreamsTable) > 0
	default:
		return false
	}
}

