package main

import (
	"fmt"
	m "math"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

// Ackley function with recommended parameters from
// https://www.sfu.ca/~ssurjano/ackley.html.
func Ackley(X []float64) float64 {
	var (
		a, b, c = 20.0, 0.2, 2 * m.Pi
		s1      float64
		s2      float64
		d       = float64(len(X))
	)
	for _, x := range X {
		s1 += x * x
		s2 += m.Cos(c * x)
	}
	return -a*m.Exp(-b*m.Sqrt(s1/d)) - m.Exp(s2/d) + a + m.Exp(1)
}

func main() {
	// Instantiate SPSO
	var de, err = eaopt.NewDefaultDiffEvo()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fix random number generation
	de.GA.RNG = rand.New(rand.NewSource(42))

	// Run minimization
	_, y, err := de.Minimize(Ackley, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output best encountered solution
	fmt.Printf("Found minimum of %.5f, the global minimum is 0\n", y)
}
