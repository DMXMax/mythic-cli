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
	for i := 0; i < len(r.dice); i++ {
		r.dice[i] = rand.Intn(3) - 1
	}
	return &r
}

func (r *Roll) Total() int {
	total := 0
	for i := 0; i < len(r.dice); i++ {
		total += r.dice[i]
	}
	for m := range r.Modifiers {
		total += int(r.Modifiers[m].Mod)
	}

	return total
}

func (r *Roll) String() string {
	return fmt.Sprintf("{ %d, %d, %d, %d } %d",
		r.dice[0], r.dice[1], r.dice[2], r.dice[3], r.Total())
}
