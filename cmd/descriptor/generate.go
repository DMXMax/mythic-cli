package descriptor

import (
	"math/rand"
	"strings"

	"github.com/DMXMax/mge/util/elements"
)

// getRandomEntry returns a random entry from the specified element table.
func getRandomEntry(tableName string) string {
	tableName = strings.ToLower(strings.TrimSpace(tableName))
	
	// Handle special cases for generic tables
	if tableName == "actions" {
		if len(elements.ActionTable1) > 0 {
			return elements.ActionTable1[rand.Intn(len(elements.ActionTable1))]
		}
		return ""
	}
	if tableName == "actions2" {
		if len(elements.ActionTable2) > 0 {
			return elements.ActionTable2[rand.Intn(len(elements.ActionTable2))]
		}
		return ""
	}
	if tableName == "descriptors" || tableName == "descriptors1" {
		if len(elements.Descriptor1) > 0 {
			return elements.Descriptor1[rand.Intn(len(elements.Descriptor1))]
		}
		return ""
	}
	if tableName == "descriptors2" {
		if len(elements.Descriptor2) > 0 {
			return elements.Descriptor2[rand.Intn(len(elements.Descriptor2))]
		}
		return ""
	}
	
	// Handle registered element tables
	switch tableName {
	case "characters":
		if len(elements.CharacterDescriptors) > 0 {
			return elements.CharacterDescriptors[rand.Intn(len(elements.CharacterDescriptors))]
		}
	case "locations":
		if len(elements.LocationTable) > 0 {
			return elements.LocationTable[rand.Intn(len(elements.LocationTable))]
		}
	case "objects":
		if len(elements.ObjectDescriptors) > 0 {
			return elements.ObjectDescriptors[rand.Intn(len(elements.ObjectDescriptors))]
		}
	case "adventure tone":
		if len(elements.AdventureToneTable) > 0 {
			return elements.AdventureToneTable[rand.Intn(len(elements.AdventureToneTable))]
		}
	case "alien species descriptors":
		if len(elements.AlienSpeciesDescriptorsTable) > 0 {
			return elements.AlienSpeciesDescriptorsTable[rand.Intn(len(elements.AlienSpeciesDescriptorsTable))]
		}
	case "animal actions":
		if len(elements.AnimalActionsTable) > 0 {
			return elements.AnimalActionsTable[rand.Intn(len(elements.AnimalActionsTable))]
		}
	case "army descriptors":
		if len(elements.ArmyDescriptorsTable) > 0 {
			return elements.ArmyDescriptorsTable[rand.Intn(len(elements.ArmyDescriptorsTable))]
		}
	case "cavern descriptors":
		if len(elements.CavernDescriptorsTable) > 0 {
			return elements.CavernDescriptorsTable[rand.Intn(len(elements.CavernDescriptorsTable))]
		}
	case "character actions combat":
		if len(elements.CharacterActionsCombatTable) > 0 {
			return elements.CharacterActionsCombatTable[rand.Intn(len(elements.CharacterActionsCombatTable))]
		}
	case "character actions general":
		if len(elements.CharacterActionsGeneralTable) > 0 {
			return elements.CharacterActionsGeneralTable[rand.Intn(len(elements.CharacterActionsGeneralTable))]
		}
	case "character appearance":
		if len(elements.CharacterAppearanceTable) > 0 {
			return elements.CharacterAppearanceTable[rand.Intn(len(elements.CharacterAppearanceTable))]
		}
	case "character background":
		if len(elements.CharacterBackgroundTable) > 0 {
			return elements.CharacterBackgroundTable[rand.Intn(len(elements.CharacterBackgroundTable))]
		}
	case "character conversations":
		if len(elements.CharacterConversationsTable) > 0 {
			return elements.CharacterConversationsTable[rand.Intn(len(elements.CharacterConversationsTable))]
		}
	case "character identity":
		if len(elements.CharacterIdentityTable) > 0 {
			return elements.CharacterIdentityTable[rand.Intn(len(elements.CharacterIdentityTable))]
		}
	case "character motivations":
		if len(elements.CharacterMotivationsTable) > 0 {
			return elements.CharacterMotivationsTable[rand.Intn(len(elements.CharacterMotivationsTable))]
		}
	case "character personality":
		if len(elements.CharacterPersonalityTable) > 0 {
			return elements.CharacterPersonalityTable[rand.Intn(len(elements.CharacterPersonalityTable))]
		}
	case "character skills":
		if len(elements.CharacterSkillsTable) > 0 {
			return elements.CharacterSkillsTable[rand.Intn(len(elements.CharacterSkillsTable))]
		}
	case "character traits flaws":
		if len(elements.CharacterTraitsFlawsTable) > 0 {
			return elements.CharacterTraitsFlawsTable[rand.Intn(len(elements.CharacterTraitsFlawsTable))]
		}
	case "city descriptors":
		if len(elements.CityDescriptorsTable) > 0 {
			return elements.CityDescriptorsTable[rand.Intn(len(elements.CityDescriptorsTable))]
		}
	case "civilization descriptors":
		if len(elements.CivilizationDescriptorsTable) > 0 {
			return elements.CivilizationDescriptorsTable[rand.Intn(len(elements.CivilizationDescriptorsTable))]
		}
	case "creature abilities":
		if len(elements.CreatureAbilitiesTable) > 0 {
			return elements.CreatureAbilitiesTable[rand.Intn(len(elements.CreatureAbilitiesTable))]
		}
	case "creature descriptors":
		if len(elements.CreatureDescriptorsTable) > 0 {
			return elements.CreatureDescriptorsTable[rand.Intn(len(elements.CreatureDescriptorsTable))]
		}
	case "cryptic message":
		if len(elements.CrypticMessageTable) > 0 {
			return elements.CrypticMessageTable[rand.Intn(len(elements.CrypticMessageTable))]
		}
	case "curses":
		if len(elements.CursesTable) > 0 {
			return elements.CursesTable[rand.Intn(len(elements.CursesTable))]
		}
	case "domicile descriptors":
		if len(elements.DomicileDescriptorsTable) > 0 {
			return elements.DomicileDescriptorsTable[rand.Intn(len(elements.DomicileDescriptorsTable))]
		}
	case "dungeon descriptors":
		if len(elements.DungeonDescriptorsTable) > 0 {
			return elements.DungeonDescriptorsTable[rand.Intn(len(elements.DungeonDescriptorsTable))]
		}
	case "dungeon traps":
		if len(elements.DungeonTrapsTable) > 0 {
			return elements.DungeonTrapsTable[rand.Intn(len(elements.DungeonTrapsTable))]
		}
	case "forest descriptors":
		if len(elements.ForestDescriptorsTable) > 0 {
			return elements.ForestDescriptorsTable[rand.Intn(len(elements.ForestDescriptorsTable))]
		}
	case "gods":
		if len(elements.GodsTable) > 0 {
			return elements.GodsTable[rand.Intn(len(elements.GodsTable))]
		}
	case "legends":
		if len(elements.LegendsTable) > 0 {
			return elements.LegendsTable[rand.Intn(len(elements.LegendsTable))]
		}
	case "magic item descriptors":
		if len(elements.MagicItemDescriptorsTable) > 0 {
			return elements.MagicItemDescriptorsTable[rand.Intn(len(elements.MagicItemDescriptorsTable))]
		}
	case "mutation descriptors":
		if len(elements.MutationDescriptorsTable) > 0 {
			return elements.MutationDescriptorsTable[rand.Intn(len(elements.MutationDescriptorsTable))]
		}
	case "names":
		if len(elements.NamesTable) > 0 {
			return elements.NamesTable[rand.Intn(len(elements.NamesTable))]
		}
	case "noble house":
		if len(elements.NobleHouseTable) > 0 {
			return elements.NobleHouseTable[rand.Intn(len(elements.NobleHouseTable))]
		}
	case "plot twists":
		if len(elements.PlotTwistsTable) > 0 {
			return elements.PlotTwistsTable[rand.Intn(len(elements.PlotTwistsTable))]
		}
	case "powers":
		if len(elements.PowersTable) > 0 {
			return elements.PowersTable[rand.Intn(len(elements.PowersTable))]
		}
	case "scavenging results":
		if len(elements.ScavengingResultsTable) > 0 {
			return elements.ScavengingResultsTable[rand.Intn(len(elements.ScavengingResultsTable))]
		}
	case "smells":
		if len(elements.SmellsTable) > 0 {
			return elements.SmellsTable[rand.Intn(len(elements.SmellsTable))]
		}
	case "sounds":
		if len(elements.SoundsTable) > 0 {
			return elements.SoundsTable[rand.Intn(len(elements.SoundsTable))]
		}
	case "spell effects":
		if len(elements.SpellEffectsTable) > 0 {
			return elements.SpellEffectsTable[rand.Intn(len(elements.SpellEffectsTable))]
		}
	case "starship descriptors":
		if len(elements.StarshipDescriptorsTable) > 0 {
			return elements.StarshipDescriptorsTable[rand.Intn(len(elements.StarshipDescriptorsTable))]
		}
	case "terrain descriptors":
		if len(elements.TerrainDescriptorsTable) > 0 {
			return elements.TerrainDescriptorsTable[rand.Intn(len(elements.TerrainDescriptorsTable))]
		}
	case "undead descriptors":
		if len(elements.UndeadDescriptorsTable) > 0 {
			return elements.UndeadDescriptorsTable[rand.Intn(len(elements.UndeadDescriptorsTable))]
		}
	case "visions dreams":
		if len(elements.VisionsDreamsTable) > 0 {
			return elements.VisionsDreamsTable[rand.Intn(len(elements.VisionsDreamsTable))]
		}
	}
	
	return ""
}

// getRandomEntries returns multiple random entries from the specified element table.
func getRandomEntries(tableName string, count int) []string {
	entries := make([]string, 0, count)
	for i := 0; i < count; i++ {
		entry := getRandomEntry(tableName)
		if entry != "" {
			entries = append(entries, entry)
		}
	}
	return entries
}

