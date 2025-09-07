# Mythic CLI

A command-line tool for solo RPG gaming using the Mythic Game Master Emulator system. Perfect for solo RPG adventures, GM-less gaming, and story generation.

## Features

- **Interactive Shell**: Command-line interface for game management
- **Dice Rolling**: Mythic fate chart integration with customizable odds and chaos factor
- **Game State Persistence**: SQLite database for saving and loading games
- **Story Logging**: Automatic logging of dice rolls and game events
- **Chaos Factor Management**: Dynamic chaos factor tracking for story complexity
- **Character and Scene Management**: Tools for managing game elements

## Installation

### Prerequisites

- Go 1.24 or later
- Git

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/DMXMax/cli-test.git
cd cli-test
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

This will give you a command prompt where you can enter various commands.

### Basic Commands

#### Game Management

- `game` or `g` - Show current game information or usage
- `game new <name> [chaos]` - Create a new game with optional chaos factor
- `game load <name>` - Load an existing game
- `game save` - Save the current game
- `game list` - List all available games
- `game chaos <value>` - Set the chaos factor (1-9)

#### Dice Rolling

- `roll` - Roll on the Mythic fate chart with current game settings
- `roll -o <odds>` - Roll with specific odds (e.g., "impossible", "very unlikely", "unlikely", "50/50", "likely", "very likely", "near sure thing", "a sure thing")
- `roll -c <chaos>` - Roll with specific chaos factor (1-9)
- `roll -o <odds> -c <chaos>` - Roll with both specific odds and chaos factor

#### Logging

- `log` - Show recent game log entries
- `log <number>` - Show last N log entries

#### Scene Management

- `scene` - Scene management commands

#### Shell Commands

- `help` - Show help for available commands
- `quit` - Exit the shell

### Example Session

```
shell> game new "My Adventure" 5
New game created: My Adventure

My Adventure> roll -o "likely"
Yes, and...

My Adventure> roll -o "unlikely" -c 7
No, but...

My Adventure> game save
Game saved successfully

My Adventure> log 5
[Recent log entries displayed]

My Adventure> quit
Goodbye!
```

## Game Mechanics

### Chaos Factor

The chaos factor (1-9) affects the likelihood of extreme results:
- **Low Chaos (1-3)**: More predictable outcomes
- **Medium Chaos (4-6)**: Balanced results
- **High Chaos (7-9)**: More extreme and unexpected outcomes

### Odds

The Mythic system uses descriptive odds:
- **Impossible** (1)
- **Very Unlikely** (2)
- **Unlikely** (3)
- **50/50** (4)
- **Likely** (5)
- **Very Likely** (6)
- **Near Sure Thing** (7)
- **A Sure Thing** (8)

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

## Development

### Project Structure

```
├── cmd/                 # Command implementations
│   ├── game/           # Game management commands
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
