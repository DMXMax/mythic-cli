package input

import (
    "bufio"
    "os"
)

// Prompter abstracts user input prompts so code can work
// both inside the liner-based shell and in non-interactive runs.
type Prompter interface {
    Prompt(prompt string) (string, error)
}

var global Prompter

// SetPrompter sets the global prompter implementation.
func SetPrompter(p Prompter) { global = p }

// Ask requests a line of input from the user. If a Prompter has been
// registered, it's used; otherwise it reads a line from stdin.
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

