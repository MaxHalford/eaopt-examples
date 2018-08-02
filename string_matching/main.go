package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/MaxHalford/eaopt"
)

var (
	corpus = strings.Split("abcdefghijklmnopqrstuvwxyz ", "")
	target = strings.Split("hello world", "")
)

// Strings is a slice of strings.
type Strings []string

// Evaluate a Strings slice by counting the number of mismatches between itself
// and the target string.
func (X Strings) Evaluate() (mismatches float64, err error) {
	for i, s := range X {
		if s != target[i] {
			mismatches++
		}
	}
	return
}

// Mutate a Strings slice by replacing it's elements by random characters
// contained in  a corpus.
func (X Strings) Mutate(rng *rand.Rand) {
	eaopt.MutUniformString(X, corpus, 3, rng)
}

// Crossover a Strings slice with another by applying 2-point crossover.
func (X Strings) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXString(X, Y.(Strings), 2, rng)
}

// MakeStrings creates random Strings slices by picking random characters from a
// corpus.
func MakeStrings(rng *rand.Rand) eaopt.Genome {
	return Strings(eaopt.InitUnifString(uint(len(target)), corpus, rng))
}

// Clone a Strings slice..
func (X Strings) Clone() eaopt.Genome {
	var XX = make(Strings, len(X))
	copy(XX, X)
	return XX
}

func main() {
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}
	ga.NGenerations = 30

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		// Concatenate the elements from the best individual and display the result
		var buffer bytes.Buffer
		for _, letter := range ga.HallOfFame[0].Genome.(Strings) {
			buffer.WriteString(letter)
		}
		fmt.Printf("%d) Result -> %s (%.0f mismatches)\n", ga.Generations, buffer.String(), ga.HallOfFame[0].Fitness)
	}

	// Run the GA
	ga.Minimize(MakeStrings)
}
