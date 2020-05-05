package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func randomRange(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Perm(size)
}

func randomArray(size, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	a := make([]int, size)
	for i := range(a) {
		a[i] = r.Intn(max)
	}
	return a
}

type SortFunc func([]int)

var algos = map[string]SortFunc {
	"Bubble sort": BubbleSort,
	"Comb sort": CombSort,
	"Cocktail sort": CocktailSort,
	"Gnome sort": GnomeSort,
	"Merge sort": MergeSort,
	"Selection sort": SelectionSort,
}

func runTest(t *testing.T, name string, fn SortFunc, a []int) {
	fn(a)
	if !sort.IntsAreSorted(a) {
		t.Error(name, "returned:", a)
	}
}

func TestSortRandom(t *testing.T) {
	for name, fn := range algos {
		a := randomArray(100, 100)
		runTest(t, name, fn, a)
	}
}

func TestSortRandomDistinct(t *testing.T) {
	for name, fn := range algos {
		a := randomRange(100)
		runTest(t, name, fn, a)
	}
}

func TestSortAllDuplicates(t *testing.T) {
	for name, fn := range algos {
		a := make([]int, 100)
		for i := range(a) {
			a[i] = 42
		}
		runTest(t, name, fn, a)
	}
}

func TestSortAlreadySorted(t *testing.T) {
	for name, fn := range algos {
		a := make([]int, 100)
		for i := range(a) {
			a[i] = i
		}
		runTest(t, name, fn, a)
	}
}

func TestSortAlreadySortedDescending(t *testing.T) {
	for name, fn := range algos {
		a := make([]int, 100)
		for i := range(a) {
			a[i] = 99 - i
		}
		runTest(t, name, fn, a)
	}
}

func TestSortEmpty(t *testing.T) {
	for name, fn := range algos {
		a := []int{}
		fn(a)
		if len(a) != 0 {
			t.Error(name, "returned:", a)
		}
	}
}
