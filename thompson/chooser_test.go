package thompson

import (
	"math/rand"
	"os"
	"testing"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	candidateN = 10
	N          = 1000
)

var (
	simulations = [N][candidateN]int{}
)

func TestMain(m *testing.M) {
	Seed(0)

	// prepare fake result for simulation
	for i := 0; i < candidateN; i++ {
		b := distuv.Bernoulli{}
		for b.P <= 0 {
			b.P = rand.Float64()
		}
		for j := 0; j < N; j++ {
			simulations[j][i] = int(b.Rand()) // 0 or 1 from Bernoulli distribution
		}
	}

	os.Exit(m.Run())
}

func TestChoose(t *testing.T) {
	candidates := make([][2]float64, candidateN)

	var thompsonHit, randomlyHit int
	for _, simulation := range simulations {
		thompsonChoose, _ := Choose(candidates)
		randomlyChoose := rand.Intn(candidateN)

		if simulation[thompsonChoose] == 1 {
			candidates[thompsonChoose][0]++
		} else {
			candidates[thompsonChoose][1]++
		}

		thompsonHit += simulation[thompsonChoose]
		randomlyHit += simulation[randomlyChoose]
	}

	if thompsonHit > randomlyHit*2 {
		// TEST PASS: Thompson Sampling should be 2x better than random choose.
	} else {
		t.Errorf("Thompson Sampling is not 2x better than random choose (thompsonHit: %d, randomlyHit: %d)", thompsonHit, randomlyHit)
	}
}
