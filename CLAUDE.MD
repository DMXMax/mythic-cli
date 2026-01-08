# mythic-cli - Interactive Command-Line Tool for Mythic GME

## Project Overview

An interactive command-line tool for solo RPG gaming using the **Mythic Game Master Emulator 2nd Edition** system. Provides a full-featured shell interface for managing games, rolling dice, tracking story logs, and generating plot points.

## Purpose

Enables solo RPG players to:
- Create and manage multiple game sessions
- Roll on the Mythic Fate Chart with customizable odds/chaos
- Roll Fate/Fudge dice (4dF) with skill/difficulty mechanics
- Track story events with automatic logging
- Manage Threads (adventure objectives) and Characters (NPCs) with weighted lists
- Generate plot points based on story themes
- Export games to Markdown via templates
- Maintain persistent game state in SQLite database

## Project Structure

```
mythic-cli/
├── cmd/                    # Command implementations
│   ├── game/              # Game management commands
│   │   ├── create.go      # Create/select games
│   │   ├── chaos.go       # Chaos factor management
│   │   ├── info.go        # Game information display
│   │   ├── list.go        # List all games
│   │   ├── load.go        # Load existing games
│   │   ├── plotpoint.go   # Plot point generation
│   │   ├── remove.go      # Delete games
│   │   ├── save.go        # Save game state
│   │   ├── export.go      # Export to Markdown
│   │   └── game.go        # Main game command
│   ├── log/               # Logging commands
│   │   ├── add.go         # Add manual log entries
│   │   ├── list.go        # View log entries
│   │   ├── remove.go      # Remove log entries
│   │   └── log.go         # Main log command
│   ├── roll/              # Dice rolling commands
│   │   ├── roll.go        # Fate Chart rolling
│   │   └── rollfate.go    # Fate dice (4dF) rolling
│   ├── scene/             # Scene management (placeholder)
│   │   └── scene.go       # Scene commands
│   └── root.go            # Root command and shell
├── util/                  # Utility packages
│   ├── db/                # Database utilities
│   │   └── db.go          # GORM setup and migrations
│   ├── dice/              # Dice rolling utilities
│   │   └── fate.go        # Fate dice implementation
│   └── game/              # Game data structures
│       └── game.go        # Game state management
├── data/                  # Data storage
│   └── templates/         # Export templates
│       └── game.md.tmpl   # Default Markdown template
├── main.go                # Application entry point
├── go.mod                 # Module definition
└── README.md              # User documentation
```

## Key Features

### Interactive Shell

- Persistent command history (`~/.mythic-cli_history`)
- Line editing with Up/Down arrow navigation
- Auto-completion support
- Contextual prompt showing current game name

### Game Management

**Commands:**
- `game create <name> [--chaos N]` - Create new or select existing game
- `game load <name>` - Load a game by name
- `game list` - List all games
- `game info` - Show current game details
- `game chaos [N]` - Get/set chaos factor (1-9)
- `game plotpoint [--verbose]` - Generate plot point from themes
- `game remove <name>` - Delete game and all logs
- `game export [-o file] [-t template] [-f]` - Export to Markdown

### Dice Rolling

**Fate Chart:**
```bash
roll                      # Default: 50/50 odds, current chaos
roll -o likely            # Specific odds
roll -c 7                 # Specific chaos
roll -o unlikely -c 3     # Both odds and chaos
roll -o ?                 # List all odds
```

**Fate Dice (4dF):**
```bash
roll rf                           # Basic 4dF roll
roll rf --skill 4                 # With skill modifier
roll rf --difficulty 3            # Against difficulty
roll rf --skill 4 --difficulty 3  # Combined
roll rf --opposed                 # Opposed roll
```

### Logging

**Commands:**
- `log` or `log print [N]` - Show last N entries (default: 20)
- `log add <message>` - Add manual story log
- `log remove [N]` - Remove last N entries

**Auto-logging:**
- Fate Chart rolls automatically log with format: `(C:5) -> likely - 42: Yes`
- Fate dice rolls log with format: `4dF { 1, 0, -1, 1 } +1; skill 3 -> 4 vs diff 2: Success (+2)`

### Database Schema

**Tables:**
- `games` - Game sessions (name, chaos, created_at, updated_at)
- `log_entries` - Story and dice logs (type, message, game_id, created_at)
- `threads` - Adventure objectives (name, description, weight 1-3, status)
- `characters` - NPCs (name, description, weight 1-3, status, notes)

**Location:** `~/.mythic-db/games.db`

## Development

### Requirements

- Go 1.24+
- SQLite 3

### Dependencies

```
github.com/spf13/cobra           # CLI framework
gorm.io/gorm                     # ORM
gorm.io/driver/sqlite            # SQLite driver
github.com/rs/zerolog            # Logging
github.com/chzyer/readline       # Interactive shell
github.com/DMXMax/mge            # Mythic GME library
```

### Building

```bash
# Build the CLI
go build -o mythic-cli

# Run the interactive shell
./mythic-cli shell
```

### Testing

```bash
# Run tests
go test ./... -v

# Test specific package
go test ./cmd/game -v
```

## Important Implementation Details

### Chaos Factor

- User-facing range: 1-9
- Default: 4
- Stored directly in database as user-facing value
- Converted to internal (0-8) when calling mge library

### Odds Normalization

The CLI normalizes odds input:
- Case-insensitive: "LIKELY" = "likely"
- Handles "50/50" format
- Accepts numeric values (0-8)
- Suggests matches on ambiguous input

### Shell Command Parsing

Commands are parsed using Cobra with custom flag handling:
- Flags must come before positional arguments
- Message text captured after flags: `roll -o likely This is my question`
- Quote multi-word odds: `roll -o "nearly certain"`

### Database Operations

**GORM Configuration:**
- Silent logger (no SQL output to users)
- Auto-migration on startup
- Timestamps automatically managed

**Duplicate Game Names:**
- `game create` loads existing game if name exists
- UNIQUE constraint on game names handled gracefully

### Export Templates

**Default Template:** `data/templates/game.md.tmpl`

**Available Helpers:**
```go
{{ formatTime .CreatedAt "2006-01-02 15:04:05" }}  // Format timestamps
{{ oddsName 5 }}                                    // Convert odds value to name
```

**Template Data:**
```go
type Game struct {
    Name       string
    Chaos      int
    CreatedAt  time.Time
    UpdatedAt  time.Time
    LogEntries []LogEntry
    Threads    []Thread
    Characters []Character
}
```

## Common Development Tasks

### Adding a New Command

1. Create file in appropriate `cmd/` subdirectory
2. Define Cobra command with flags and args
3. Register with parent command in package init
4. Add help text and examples
5. Update README.md

Example:
```go
var myCmd = &cobra.Command{
    Use:   "my-command [args]",
    Short: "Brief description",
    Long:  "Detailed description",
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation
    },
}

func init() {
    parentCmd.AddCommand(myCmd)
    myCmd.Flags().StringP("flag", "f", "default", "description")
}
```

### Adding a New Database Table

1. Define model in `util/game/game.go` or create new file
2. Add GORM struct tags
3. Add to auto-migration in `util/db/db.go`
4. Test migration with fresh database

### Modifying Roll Behavior

**Fate Chart rolls:** Edit `cmd/roll/roll.go`
**Fate dice rolls:** Edit `cmd/roll/rollfate.go`
**Dice library:** Edit `util/dice/fate.go` or upstream to `mge`

## Code Style Guidelines

- Use structured logging (zerolog) for debugging
- Keep handlers focused (parse flags → validate → execute → output)
- Validate user input early
- Provide clear error messages
- Follow Cobra conventions for command structure
- Use GORM best practices (preload associations, check errors)

## Configuration Files

### Shell History

**Location:** `~/.mythic-cli_history`
**Format:** Plain text, one command per line
**Management:** Automatically maintained by readline

### Database

**Location:** `~/.mythic-db/games.db`
**Backup:** Copy file to backup location
**Reset:** Delete file to start fresh (WARNING: deletes all games)

## Debugging

### Enable Verbose Logging

Set log level in `main.go`:
```go
zerolog.SetGlobalLevel(zerolog.DebugLevel)
```

### Database Inspection

```bash
# Open database with sqlite3
sqlite3 ~/.mythic-db/games.db

# View tables
.tables

# Query games
SELECT * FROM games;

# Query logs for a game
SELECT * FROM log_entries WHERE game_id = '...';
```

## Integration with mge Library

### Chaos Conversion

```go
// User input (1-9) -> Internal (0-8)
internalChaos := chart.ChaosUserToInternal(game.Chaos)

// Call mge
result := chart.FateChart.RollOdds(odds, internalChaos)
```

### Odds Matching

```go
// Match by prefix
matches := chart.MatchOddsPrefix(input)
if len(matches) != 1 {
    // Handle ambiguous or no match
}
odds := chart.Odds(matches[0])
```

## Known Issues & Fixes

### Flag Parsing Order

Flags must come before positional args:
- ✅ `roll -o likely -c 5 My question`
- ❌ `roll My question -o likely -c 5`

### Database Constraints

Games with duplicate names handled gracefully:
- Creates new game if name doesn't exist
- Loads existing game if name exists
- No error thrown

### GORM Silent Logs

Database logs hidden from users for cleaner output:
```go
db.Logger = logger.Default.LogMode(logger.Silent)
```

## Legal & Licensing

**Important:** Implements copyrighted game mechanics from **Mythic Game Master Emulator 2nd Edition** by Tana Pigeon and Word Mill Games.

- Users must own a legitimate copy of the official product
- Contact Word Mill Games for licensing: www.wordmillgames.com

## Quick Reference

### Most Common Commands

```bash
game create "My Adventure"        # Start a game
roll                              # Roll fate chart
roll rf                           # Roll fate dice
log                               # View recent logs
game info                         # Show game details
game plotpoint                    # Generate plot point
quit                              # Exit shell
```

### Game State Management

```bash
game create <name> --chaos 5      # Create with custom chaos
game load <name>                  # Switch games
game chaos 7                      # Update chaos
game list                         # See all games
game remove <name>                # Delete game
```

### Export

```bash
game export                       # Export current game
game export "Other Game"          # Export specific game
game export -o output.md -f       # Force overwrite
game export -t custom.tmpl        # Custom template
```

## Notes for AI Assistants

When working on this codebase:
1. **Maintain shell UX** - Commands should be intuitive and forgiving
2. **Test with real database** - SQLite behavior differs from mocks
3. **Respect mge library API** - Coordinate breaking changes
4. **Keep output clean** - Users don't want verbose logs
5. **Validate chaos range** - Must be 1-9 user-facing
6. **Auto-log appropriately** - Dice rolls yes, info commands no
7. **Preserve history** - Shell history is important UX feature
8. **Handle edge cases** - Empty games, missing files, etc.
