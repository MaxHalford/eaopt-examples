package main

import (
	"fmt"
	"image/gif"
	"math"
	"math/rand"
	"os"

	"github.com/MaxHalford/eaopt"
)

const (
	// GRIDSIZE determines the size of the grid, ie there will be GRIDSIZE^2 points
	GRIDSIZE = 5
	// GENERATIONS determines how many generations the GA will be run for
	GENERATIONS = 800
)

// A Point is defined with (x, y) coordinates.
type Point struct {
	x, y float64
}

// A Path is a slice of Points.
type Path []Point

// At method from Slice
func (p Path) At(i int) interface{} {
	return p[i]
}

// Set method from Slice
func (p Path) Set(i int, v interface{}) {
	p[i] = v.(Point)
}

// Len method from Slice
func (p Path) Len() int {
	return len(p)
}

// Swap method from Slice
func (p Path) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Slice method from Slice
func (p Path) Slice(a, b int) eaopt.Slice {
	return p[a:b]
}

// Split method from Slice
func (p Path) Split(k int) (eaopt.Slice, eaopt.Slice) {
	return p[:k], p[k:]
}

// Append method from Slice
func (p Path) Append(q eaopt.Slice) eaopt.Slice {
	return append(p, q.(Path)...)
}

// Replace method from Slice
func (p Path) Replace(q eaopt.Slice) {
	copy(p, q.(Path))
}

// Copy method from Slice
func (p Path) Copy() eaopt.Slice {
	var clone = make(Path, len(p))
	copy(clone, p)
	return clone
}

// Evaluate a Path by summing the consecutive Euclidean distances.
func (p Path) Evaluate() (distance float64, err error) {
	for i := 0; i < len(p)-1; i++ {
		distance += math.Sqrt(math.Pow(p[i+1].x-p[i].x, 2) + math.Pow(p[i+1].y-p[i].y, 2))
	}
	return
}

// Mutate a Path by applying by permutation mutation and/or splice mutation.
func (p Path) Mutate(rng *rand.Rand) {
	if rng.Float64() < 0.35 {
		eaopt.MutPermute(p, 3, rng)
	}
	if rng.Float64() < 0.45 {
		eaopt.MutSplice(p, rng)
	}
}

// Crossover a Path with another Path by using Partially Mixed Crossover (PMX).
func (p Path) Crossover(q eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossPMX(p, q.(Path), rng)
}

// Clone a Path.
func (p Path) Clone() eaopt.Genome {
	var clone = make(Path, len(p))
	copy(clone, p)
	return clone
}

// MakePath creates a slice of Points along a grid and then shuffles the slice.
func MakePath(rng *rand.Rand) eaopt.Genome {
	var (
		path = make(Path, GRIDSIZE*GRIDSIZE)
		p    = 0
	)
	for i := 0; i < GRIDSIZE; i++ {
		for j := 0; j < GRIDSIZE; j++ {
			path[p] = Point{float64(i), float64(j)}
			p++
		}
	}
	// Shuffle the points
	for i := range path {
		var j = rng.Intn(i + 1)
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func main() {
	var (
		ga, err = eaopt.NewDefaultGAConfig().NewGA()
		outGif  = &gif.GIF{}
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		// Store the drawing for the current best path
		var img = drawPath(ga.HallOfFame[0].Genome.(Path), ga.Generations, ga.HallOfFame[0].Fitness)
		outGif.Image = append(outGif.Image, img)
		outGif.Delay = append(outGif.Delay, 0)
	}

	// Run the GA
	ga.Minimize(MakePath)

	// Print the best obtained solution vs. the optimal solution
	var optimal = float64((GRIDSIZE + 1) * (GRIDSIZE - 1))
	fmt.Printf("Obtained %f\n", ga.HallOfFame[0].Fitness)
	fmt.Printf("Optimal is %d\n", int(optimal))
	fmt.Printf("Off by %f percent\n", 100*(ga.HallOfFame[0].Fitness-optimal)/optimal)

	// Save to out.gif
	var outFile, _ = os.OpenFile("progress.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer outFile.Close()
	gif.EncodeAll(outFile, outGif)
}
