package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/MaxHalford/gago"
)

var (
	corpus = strings.Split("abcdefghijklmnopqrstuvwxyz ", "")
	target = strings.Split("hello world", "")
)

// Strings is a slice of strings.
type Strings []string

// Evaluate a Strings slice by counting the number of mismatches between itself
// and the target string.
func (X Strings) Evaluate() (mismatches float64) {
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
	gago.MutUniformString(X, corpus, 3, rng)
}

// Crossover a Strings slice with another by applying 2-point crossover.
func (X Strings) Crossover(Y gago.Genome, rng *rand.Rand) (gago.Genome, gago.Genome) {
	var o1, o2 = gago.CrossGNXString(X, Y.(Strings), 2, rng)
	return Strings(o1), Strings(o2)
}

// MakeStrings creates random Strings slices by picking random characters from a
// corpus.
func MakeStrings(rng *rand.Rand) gago.Genome {
	return Strings(gago.InitUnifString(len(target), corpus, rng))
}

// Clone a Strings slice..
func (X Strings) Clone() gago.Genome {
	var XX = make(Strings, len(X))
	copy(XX, X)
	return XX
}

func main() {
	var ga = gago.Generational(MakeStrings)
	ga.Initialize()

	for i := 1; i < 30; i++ {
		ga.Enhance()
		// Concatenate the elements from the best individual and display the result
		var buffer bytes.Buffer
		for _, letter := range ga.Best.Genome.(Strings) {
			buffer.WriteString(letter)
		}
		fmt.Printf("Result -> %s (%.0f mismatches)\n", buffer.String(), ga.Best.Fitness)
	}
}
