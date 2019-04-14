package thompson

import (
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

var src rand.Source

func Seed(s int64) {
	src = rand.NewSource(uint64(s))
}

func Choose(candidates [][2]float64) (selected int, maxValue float64) {
	for i, candidate := range candidates {
		a := math.Max(candidate[0], 1)
		b := math.Max(candidate[1], 1)

		value := distuv.Beta{Alpha: a, Beta: b, Src: src}.Rand()
		if value > maxValue {
			maxValue = value
			selected = i
		}
	}
	return selected, maxValue
}
