package dice

import (
	"fmt"
	"math/rand"
)

// RollModifier is a struct that holds the modifier and a description of the modifier
type RollModifier struct {
	Mod         int8
	Description string
}

// Roll is a struct that holds the dice and a description of the roll
type Roll struct {
	dice        [4]int
	Description string
	Modifiers   []RollModifier
}

func RollFate() *Roll {
	r := Roll{}
	for i := range r.dice {
		r.dice[i] = rand.Intn(3) - 1
	}
	return &r
}

// DiceTotal returns the sum of the dice without any modifiers.
func (r *Roll) DiceTotal() int {
	total := 0
	for _, die := range r.dice {
		total += die
	}
	return total
}

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

func (r *Roll) String() string {
	return fmt.Sprintf("{ %d, %d, %d, %d } %+d",
		r.dice[0], r.dice[1], r.dice[2], r.dice[3], r.DiceTotal())
}
