package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/MaxHalford/eaopt"
)

// NQUEENS is the size of each genome.
const NQUEENS = 15

// Positions is a slice of ints.
type Positions []int

// String prints a chess board and marks with an x the queen positions.
func (P Positions) String() string {
	var board bytes.Buffer
	for _, p := range P {
		board.WriteString(strings.Repeat(" .", p))
		board.WriteString(" \u2655")
		board.WriteString(strings.Repeat(" .", NQUEENS-p-1))
		board.WriteString("\n")
	}
	return board.String()
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Evaluate a slice of Positions by counting the number of diagonal collisions.
// Queens are on the same diagonal if there row distance is equal to their
// column distance.
func (P Positions) Evaluate() (float64, error) {
	var collisions float64
	for i := 0; i < len(P); i++ {
		for j := i + 1; j < len(P); j++ {
			if j-i == absInt(P[i]-P[j]) {
				collisions++
			}
		}
	}
	return collisions, nil
}

// Mutate a slice of Positions by permuting it's values.
func (P Positions) Mutate(rng *rand.Rand) {
	eaopt.MutPermuteInt(P, 3, rng)
}

// Crossover a slice of Positions with another by applying partially mapped
// crossover.
func (P Positions) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossPMXInt(P, Y.(Positions), rng)
}

// Clone a slice of Positions.
func (P Positions) Clone() eaopt.Genome {
	var PP = make(Positions, len(P))
	copy(PP, P)
	return PP
}

// MakeBoard creates a random slices of positions by generating random number
// permutations in [0, NQUEENS).
func MakeBoard(rng *rand.Rand) eaopt.Genome {
	var positions = make(Positions, NQUEENS)
	for i, position := range rng.Perm(NQUEENS) {
		positions[i] = position
	}
	return eaopt.Genome(positions)
}

func main() {
	var conf = eaopt.NewDefaultGAConfig()
	conf.NGenerations = 10e9 // We should stop earlier than this
	var ga, err = conf.NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add a callback to stop when the problem is solved
	ga.EarlyStop = func(ga *eaopt.GA) bool {
		return ga.HallOfFame[0].Fitness == 0
	}

	// Run the GA
	ga.Minimize(MakeBoard)

	// Display result
	fmt.Println(ga.HallOfFame[0].Genome)
	fmt.Printf("Optimal solution obtained after %d generations in %s\n", ga.Generations, ga.Age)
}
