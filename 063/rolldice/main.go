// Package rolldice contains a function that simulates a dice roll.
package rolldice

import (
	"math/rand"
	"time"
)

// RollDice returns a random number between 1 and 6 both included
func RollDice() int {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)
	return 1 + randGenerator.Intn(6)
}
