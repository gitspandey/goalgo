package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func randomArray(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Perm(size)
}

type SortFunc func([]int)

var algos = map[string]SortFunc {
	"Bubble sort": BubbleSort,
	"Comb sort": CombSort,
	"Cocktail sort": CocktailSort,
	"Gnome sort": GnomeSort,
	"Selection sort": SelectionSort,
}

func TestSort(t *testing.T) {
	for name, fn := range algos {
		a := randomArray(100)
		fn(a)
		if !sort.IntsAreSorted(a) {
			t.Error(name, "returned:", a)
		}
	}
}
