package main

import (
	"fmt"
	m "math"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

// A Vector contains float64s.
type Vector []float64

// Evaluate a Vector with the Drop-Wave function which takes two variables as
// input and reaches a minimum of -1 in X = (0, 0).
func (X Vector) Evaluate() (float64, error) {
	var (
		numerator   = 1 + m.Cos(12*m.Sqrt(m.Pow(X[0], 2)+m.Pow(X[1], 2)))
		denominator = 0.5*(m.Pow(X[0], 2)+m.Pow(X[1], 2)) + 2
	)
	return -numerator / denominator, nil
}

// Mutate a Vector by applying by resampling each element from a normal
// distribution with probability 0.8.
func (X Vector) Mutate(rng *rand.Rand) {
	eaopt.MutNormalFloat64(X, 0.8, rng)
}

// Crossover a Vector with another Vector by applying uniform crossover.
func (X Vector) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossUniformFloat64(X, Y.(Vector), rng)
}

// Clone a Vector.
func (X Vector) Clone() eaopt.Genome {
	var XX = make(Vector, len(X))
	copy(XX, X)
	return XX
}

// MakeVector returns a random vector by generating 2 values uniformally
// distributed between -10 and 10.
func MakeVector(rng *rand.Rand) eaopt.Genome {
	return Vector(eaopt.InitUnifFloat64(2, -10, 10, rng))
}

func main() {
	var conf = eaopt.NewDefaultGAConfig()
	conf.NPops = 1
	var ga, err = conf.NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Run the GA
	ga.Minimize(MakeVector)
}
