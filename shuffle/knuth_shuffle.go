package shuffle

import (
	"math/rand"
	"time"
)

// Reference: https://en.wikipedia.org/wiki/Knuth_shuffle

// Allow specifying an optional random number generator for testing
type KnuthShuffler struct {
	// A function that returns a random number >= 0 and < n
	// Equivalent to rand.Intn
	RandIntn func(int) int
}

func (shuffler *KnuthShuffler) Shuffle(a []int) {
	randIntn := shuffler.RandIntn
	if randIntn == nil {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randIntn = func(n int) int {
			return r.Intn(n)
		}
	}

	for i := len(a) - 1; i > 0; i-- {
		j := randIntn(i + 1)
		a[j], a[i] = a[i], a[j]
	}
}
