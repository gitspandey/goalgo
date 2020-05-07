package array

import (
	"math"
	"testing"
)

type BinarySearchTest testing.T

func (t *BinarySearchTest) assertBinarySearchResult(n, actual, expected, callCount, callCountLimit int) {
	if actual != expected {
		t.Fatalf("Search for %v yielded %v (expected: %v)", n, actual, expected)
	}

	if callCount > callCountLimit {
		t.Fatalf("Too many comparisons while searching for %v: %v (max expected: %v)", n, callCount, callCountLimit)
	}
}

func TestBinarySearch(t *testing.T) {
	a := make([]int, 100)
	for i := range a {
		a[i] = i
	}

	// In the worst case, binary search should make floor(log2(n)) + 1 comparisons,
	// where n is the size of the array
	callCountLimit := int(math.Floor(math.Log2(float64(len(a))))) + 1

	for _, n := range a {
		callCount := 0
		searcher := BinarySearcher{
			ElementAt: func(a []int, i int) int {
				callCount++
				return a[i]
			},
		}
		result := searcher.Search(a, n)
		t1 := BinarySearchTest(*t)
		t1.assertBinarySearchResult(n, result, n, callCount, callCountLimit)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	a := make([]int, 100)
	for i := range a {
		a[i] = 2 * i
	}

	// In the worst case, binary search should make floor(log2(n)) + 1 comparisons,
	// where n is the size of the array
	callCountLimit := int(math.Floor(math.Log2(float64(len(a))))) + 1

	for _, n := range a {
		callCount := 0
		searcher := BinarySearcher{
			ElementAt: func(a []int, i int) int {
				callCount++
				return a[i]
			},
		}
		result := searcher.Search(a, n+1)
		t1 := BinarySearchTest(*t)
		t1.assertBinarySearchResult(n+1, result, -1, callCount, callCountLimit)
	}
}

func TestBinarySearchEmptyArray(t *testing.T) {
	a := []int{}

	callCount := 0
	searcher := BinarySearcher{
		ElementAt: func(a []int, i int) int {
			callCount++
			return a[i]
		},
	}
	result := searcher.Search(a, 42)
	t1 := BinarySearchTest(*t)
	t1.assertBinarySearchResult(42, result, -1, callCount, 0)
}
