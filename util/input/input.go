// Package input provides an abstraction layer for user input prompts.
// It allows code to work both inside the liner-based interactive shell
// and in non-interactive command-line contexts.
package input

import (
	"bufio"
	"os"
)

// Prompter is an interface for prompting user input.
// Implementations can provide different input mechanisms (e.g., liner for interactive shells,
// stdin for non-interactive contexts).
type Prompter interface {
	Prompt(prompt string) (string, error)
}

var global Prompter

// SetPrompter sets the global prompter implementation.
// This is typically called during shell initialization to use liner for interactive prompts.
func SetPrompter(p Prompter) {
	global = p
}

// Ask requests a line of input from the user.
// If a Prompter has been registered via SetPrompter, it is used.
// Otherwise, it falls back to reading a line from stdin.
//
// Parameters:
//   - prompt: The prompt string to display to the user
//
// Returns the user's input line and any error that occurred.
func Ask(prompt string) (string, error) {
	if global != nil {
		return global.Prompt(prompt)
	}
	// Fallback simple prompt
	if _, err := os.Stdout.WriteString(prompt); err != nil {
		return "", err
	}
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	return line, err
}

