// Package dice provides dice rolling utilities for Fate/Fudge dice (4dF).
// Fate dice have three faces: -1, 0, and +1, resulting in a range from -4 to +4.
package dice

import (
	"fmt"
	"math/rand"
)

// RollModifier represents a modifier applied to a dice roll.
// It contains both the numeric modifier value and a description of what it represents.
type RollModifier struct {
	Mod         int8  // The numeric modifier value
	Description string // A description of what this modifier represents (e.g., "skill", "bonus")
}

// Roll represents a Fate/Fudge dice roll (4dF).
// Each die can be -1, 0, or +1, giving a total range of -4 to +4.
type Roll struct {
	dice        [4]int         // The four dice values (-1, 0, or +1)
	Description string         // Optional description of the roll
	Modifiers   []RollModifier // List of modifiers to apply to the roll
}

// RollFate rolls four Fate/Fudge dice and returns a new Roll instance.
// Each die has an equal chance of being -1, 0, or +1.
func RollFate() *Roll {
	r := Roll{}
	for i := range r.dice {
		r.dice[i] = rand.Intn(3) - 1
	}
	return &r
}

// DiceTotal returns the sum of the four dice without any modifiers.
// This represents the raw dice result before applying modifiers.
func (r *Roll) DiceTotal() int {
	total := 0
	for _, die := range r.dice {
		total += die
	}
	return total
}

// Total returns the sum of the dice plus all modifiers.
// This is the final result of the roll after all modifiers are applied.
func (r *Roll) Total() int {
	total := 0
	for _, die := range r.dice {
		total += die
	}
	for _, modifier := range r.Modifiers {
		total += int(modifier.Mod)
	}

	return total
}

// String returns a string representation of the roll.
// Format: "{ d1, d2, d3, d4 } +total" where total is the dice total without modifiers.
func (r *Roll) String() string {
	return fmt.Sprintf("{ %d, %d, %d, %d } %+d",
		r.dice[0], r.dice[1], r.dice[2], r.dice[3], r.DiceTotal())
}
