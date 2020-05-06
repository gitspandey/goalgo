package shuffle

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestKnuthShuffle(t *testing.T) {
	expected := []int{52, 6, 62, 56, 1, 87, 13, 48, 83, 2, 3, 37, 33, 61, 39, 10, 91, 7, 63, 86, 57, 41, 17, 47, 97, 22, 0, 14, 93, 55, 71, 81, 26, 43, 60, 27, 76, 74, 92, 66, 25, 98, 96, 34, 50, 49, 59, 88, 12, 28, 75, 45, 90, 30, 70, 77, 32, 94, 73, 35, 84, 18, 78, 54, 23, 85, 24, 95, 99, 79, 69, 67, 20, 72, 11, 36, 21, 16, 46, 42, 29, 82, 51, 80, 58, 68, 89, 15, 65, 19, 38, 8, 44, 9, 40, 31, 64, 4, 53, 5}

	for i := 0; i < 100; i++ {
		a := make([]int, 100)
		for i := range a {
			a[i] = i
		}

		// Fix the random number generator for testing purposes
		r := rand.New(rand.NewSource(42))
		shuffler := KnuthShuffler{
			RandIntn: func(n int) int {
				return r.Intn(n)
			},
		}
		shuffler.Shuffle(a)

		if len(a) != len(expected) {
			t.Fatalf("Shuffled array is of length: %v. Expected: %v", len(a), len(expected))
		}

		for i := range a {
			if a[i] != expected[i] {
				t.Fatalf("Mismatch on index %v: %v != %v", i, a[i], expected[i])
			}
		}
	}
}

func TestKnuthShuffleNoRepeatedShuffles(t *testing.T) {
	arrayToString := func(a []int) string {
		s := make([]string, len(a))
		for i := range a {
			s[i] = strconv.Itoa(a[i])
		}
		return strings.Join(s, ",")
	}

	// keep track of the shuffle results in a map so we can find repetitions
	shuffles := make(map[string]int)
	for i := 0; i < 10000; i++ {
		a := make([]int, 100)
		for i := range a {
			a[i] = i
		}

		shuffler := KnuthShuffler{}
		shuffler.Shuffle(a)

		if len(a) != 100 {
			t.Fatalf("Shuffled array is of length: %v. Expected: %v", len(a), 100)
		}

		key := arrayToString(a)
		freq := shuffles[key]
		if freq > 2 {
			t.Fatalf("Same shuffle result seen more than %v times", freq)
		}
		shuffles[key] = freq + 1
	}
}
