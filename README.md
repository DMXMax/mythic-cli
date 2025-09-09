# Mythic CLI

A command-line tool for solo RPG gaming using the Mythic Game Master Emulator system. Perfect for solo RPG adventures, GM-less gaming, and story generation.

## Features

- **Interactive Shell**: Command-line interface for game management
- **Dice Rolling**: Mythic fate chart integration with customizable odds and chaos factor
- **Game State Persistence**: SQLite database for saving and loading games
- **Story Logging**: Automatic logging of dice rolls and game events
- **Chaos Factor Management**: Dynamic chaos factor tracking for story complexity
- **Character and Scene Management**: Tools for managing game elements
- **Markdown Export**: Export a game and its log via a customizable template
- **Robust Error Handling**: Comprehensive validation and user feedback
- **Flexible Game Creation**: Multiple ways to create games with custom chaos factors and default odds

## Installation

### Prerequisites

- Go 1.24 or later
- Git

### Build from Source

1. Clone the repository:
```bash

git clone https://github.com/DMXMax/mythic-cli.git
cd mythic-cli
```

2. Build the application:
```bash
go build -o mythic-cli
```

3. Run the application:
```bash
./mythic-cli
```

## Usage

### Starting the Interactive Shell

To begin using the Mythic CLI, start the interactive shell:

```bash
./mythic-cli shell
```

This opens a command prompt with line editing and command history.
Tips:
- Up/Down arrows navigate history (persisted at `~/.mythic-cli_history`).
- Press Ctrl-C or Ctrl-D to exit; or type `quit`.

### Basic Commands

#### Game Management

- `game` or `g` - Show current game information or usage
- `game create <name> [--chaos <value>]` - Create a new game with optional chaos factor
- `game create <name> [-x <value>]` - Create a new game with short chaos flag
- `game new <name>` - Alias for `game create` (backward compatibility)
- `game create <name> [--odds <value>]` - Create a new game with optional default odds
- `game load <name>` - Load an existing game
- `game save` - Save the current game
- `game list` - List all available games
- `game chaos <value>` - Set the chaos factor (1-9)
- `game remove <name>` - Remove a game and all of its log entries
- `game export [name] [-o <file>] [-t <template>] [-f]` - Export current or named game to Markdown using a template (see Export section)

#### Dice Rolling

- `roll` - Roll on the Mythic fate chart with current game settings (default: likely odds, game's chaos factor)
- `roll -o <odds>` - Roll with specific odds (case-insensitive; accepts names like "nearly certain" or formats like `50/50`)
- `roll -o ?` - List all available odds names and their numeric values
- `roll -c <chaos>` - Roll with specific chaos factor (1-9) and default odds
- `roll -o <odds> -c <chaos>` - Roll with both specific odds and chaos factor
- `roll --help` - Show detailed help for the roll command

Odds input notes:
- Quote multi-word odds: `-o "nearly certain"` (or use the numeric value, e.g., `-o 7`).
- `-o 50/50` works without quotes and is normalized to "fifty fifty".
- On ambiguous input, the CLI suggests close matches without consuming your message text.

#### Logging

- `log print` - Show recent game log entries
- `log print <number>` - Show last N log entries
- `log remove <number>` - Remove last N log entries
- `log add <message>` - Add a log entry
- `log remove` - Remove all log entries
- `log --help` - Show detailed help for the log command
- `log` same as `log print`



#### Scene Management

- `scene` - Scene management commands

#### Shell Commands

- `help` - Show help for available commands
- `quit` - Exit the shell

### Example Session

```
shell> game create "My Adventure" --chaos 5
Created new game: My Adventure (Chaos: 5)

My Adventure> roll
likely - 42: Yes

My Adventure> roll -o "likely"
likely - 67: No

My Adventure> roll -c 8
likely - 15: Exceptional Yes

My Adventure> roll -o "unlikely" -c 2
unlikely - 23: Yes

My Adventure> game create "My Adventure"
Selected existing game: My Adventure (Chaos: 5)

My Adventure> gl print
likely - 42: Yes
likely - 67: No
likely - 15: Exceptional Yes
unlikely - 23: Yes

My Adventure> quit
Goodbye!
```

### Alternative Command Examples

```
# Create a game with default chaos (4)
shell> game create "Quick Adventure"

# Create a game with custom chaos using short flag
shell> game create "Chaotic Quest" -x 8

# Create a game with custom chaos using long flag
shell> game create "Epic Journey" --chaos 6

# The 'new' alias still works for backward compatibility
shell> game new "Legacy Game"

# Export the current game to Markdown (prompts if file exists)
shell> game export

# Export a named game to a specific file path without prompting
shell> game export "My Adventure" -o exports/my-adventure.md -f

# Use a custom template for export
shell> game export -t data/templates/game.md.tmpl
```

## Game Mechanics

### Chaos Factor

The chaos factor (1-9) affects the likelihood of extreme results:
- **Low Chaos (1-3)**: More predictable outcomes
- **Medium Chaos (4-6)**: Balanced results
- **High Chaos (7-9)**: More extreme and unexpected outcomes

### Odds

The Mythic system uses descriptive odds:
- **Impossible** (0)
- **Nearly Impossible** (1)
- **Very Unlikely** (2)
- **Unlikely** (3)
- **Fifty Fifty** (4)
- **Likely** (5)
- **Very Likely** (6)
- **Nearly Certain** (7)
- **Certain** (8)

### Fate Chart Results

Rolls can result in:
- **Yes, and...** - Yes with additional benefits
- **Yes** - Simple yes
- **Yes, but...** - Yes with complications
- **No, but...** - No with some benefit
- **No** - Simple no
- **No, and...** - No with additional complications

## Data Storage

Games are automatically saved to a SQLite database (`data/games.db`) with the following information:
- Game name and metadata
- Current chaos factor
- Complete log of all dice rolls and events
- Timestamps for all entries

### Game Management

- **Automatic Game Loading**: If you try to create a game with a name that already exists, the system will automatically load the existing game instead of creating a duplicate
- **Persistent State**: Games are automatically saved when you make dice rolls or add log entries
- **Database Integrity**: UNIQUE constraint errors are handled gracefully

## Recent Improvements

### Version Updates

- **Enhanced Game Creation**: The `game new` command has been renamed to `game create` for better clarity
- **Improved Error Handling**: Added comprehensive validation for chaos factors and game parameters
- **Better User Feedback**: Clear confirmation messages when creating or selecting games
- **Default Odds Setting**: Games can now be created with a default odds value
- **Flexible Flag Options**: Support for both long (`--chaos`) and short (`-x`) flags
- **Backward Compatibility**: The `new` alias still works for existing users
- **Robust Validation**: Chaos factor validation using proper range checking (1-9)
- **Fixed Help Commands**: All help commands now work properly (`--help` flag support)
- **Fixed Roll Command**: Dice rolling now works correctly with all flag combinations
- **Database Constraint Fixes**: Games with duplicate names are now handled gracefully
- **Improved User Experience**: Removed verbose trace and debug messages for cleaner output
- **Interactive Shell History**: Up/Down arrow-key history with persistent storage at `~/.mythic-cli_history`
- **Roll Odds Helper**: `-o ?` prints available odds and their indices
- **Normalized Odds Input**: Case-insensitive matching; accepts common forms like `50/50`; clearer suggestions on ambiguity
- **Game Export**: Export games to Markdown using a template with safe overwrite prompts and a `--force` option

### Command Improvements

- Fixed flag parsing order for better reliability
- Added structured logging for better debugging
- Improved error messages with proper context
- Enhanced command documentation and help text
- **Fixed Shell Flag Handling**: Commands now properly parse flags and handle default values
- **Fixed Database Issues**: UNIQUE constraint errors resolved for duplicate game names
- **Improved Roll Command**: Fixed chaos factor and odds parsing issues
- **Roll UX**: Odds parsing no longer steals the first word of your message
- **Enhanced Shell Stability**: Multiple commands in sequence now work correctly
- **Cleaner Output**: Removed verbose logging messages for better user experience
- **Silent Database Operations**: GORM database logs are now hidden from users

## Exporting to Markdown

Use `game export` to render a game and its log to a Markdown file via a Go text/template.

- Default template: `data/templates/game.md.tmpl`
- Default output file: `<game>.md`
- Overwrite behavior: If the output file exists, the CLI prompts before overwriting. Use `-f/--force` to overwrite without prompting.

Common flags:
- `-o, --out <file>`: Output path (e.g., `exports/mygame.md`)
- `-t, --template <path>`: Template file path
- `-f, --force`: Overwrite existing output without prompting

Template helpers available:
- `formatTime .CreatedAt "2006-01-02 15:04:05"` – format timestamps
- `oddsName .Odds` – turn the numeric odds into a name (e.g., "likely")

## Development

### Project Structure

```
├── cmd/                 # Command implementations
│   ├── game/           # Game management commands
│   │   ├── create.go   # Game creation command (formerly new)
│   │   ├── chaos.go    # Chaos factor management
│   │   ├── load.go     # Game loading and listing
│   │   ├── save.go     # Game saving
│   │   └── game.go     # Main game command
│   ├── log/            # Logging commands
│   ├── roll/           # Dice rolling commands
│   ├── scene/          # Scene management commands
│   └── root.go         # Root command and shell
├── util/               # Utility packages
│   ├── db/             # Database utilities
│   ├── dice/           # Dice rolling utilities
│   └── game/           # Game data structures
├── data/               # Data storage
└── main.go             # Application entry point
```

### Dependencies

- **Cobra**: CLI framework
- **GORM**: ORM for database operations
- **SQLite**: Database storage
- **Zerolog**: Logging
- **MGE**: Mythic Game Master Emulator library

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the terms specified in the LICENSE file.

## Acknowledgments

- Based on the Mythic Game Master Emulator by Tana Pigeon
- Built with the Go programming language
- Uses the Cobra CLI framework
