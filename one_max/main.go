// See http://tracer.lcc.uma.es/problems/onemax/onemax.html for a description
// of the problem.
package main

import (
	"fmt"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

// N is the size of each genome.
const N = 20

// Digits is a slice of ints.
type Digits []int

// Evaluate a slice of Digits by summing the number of 1s.
func (X Digits) Evaluate() (float64, error) {
	var sum int
	for _, d := range X {
		sum += d
	}
	return N - float64(sum), nil // We want to minimize the fitness, hence the reversing
}

// Mutate a slice of Digits by permuting it's values.
func (X Digits) Mutate(rng *rand.Rand) {
	eaopt.MutPermuteInt(X, 3, rng)
}

// Crossover a slice of Digits with another by applying 2-point crossover.
func (X Digits) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXInt(X, Y.(Digits), 2, rng)
}

// Clone a slice of Digits.
func (X Digits) Clone() eaopt.Genome {
	var XX = make(Digits, len(X))
	copy(XX, X)
	return XX
}

// MakeDigits creates a random slice of Digits by randomly picking 1s and 0s.
func MakeDigits(rng *rand.Rand) eaopt.Genome {
	var digits = make(Digits, N)
	for i := range digits {
		if rng.Float64() < 0.5 {
			digits[i] = 1
		}
	}
	return eaopt.Genome(digits)
}

func main() {
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("%d) Best fitness -> %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Add a callback to stop when the problem is solved
	ga.EarlyStop = func(ga *eaopt.GA) bool {
		return ga.HallOfFame[0].Fitness == 0
	}

	// Run the GA
	ga.Minimize(MakeDigits)
}
