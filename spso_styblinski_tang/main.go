package main

import (
	"fmt"
	m "math"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

func StyblinskiTang(X []float64) (y float64) {
	for _, x := range X {
		y += m.Pow(x, 4) - 16*m.Pow(x, 2) + 5*x
	}
	return 0.5 * y
}

func main() {
	// Instantiate SPSO
	var spso, err = eaopt.NewDefaultSPSO()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fix random number generation
	spso.GA.RNG = rand.New(rand.NewSource(42))

	// Run minimization
	_, y, err := spso.Minimize(StyblinskiTang, 2) // Dimension is 4
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output best encountered solution
	fmt.Printf("Found minimum of %.5f, the global minimum is %.5f\n", y, -39.16599*2)
}
