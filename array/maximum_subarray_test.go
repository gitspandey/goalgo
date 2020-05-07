package array

import "testing"

func assertSlicesEqual(t *testing.T, actual, expected []int) {
	if len(actual) != len(expected) {
		t.Fatalf("Array is of length: %v. Expected: %v", len(actual), len(expected))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Fatalf("Mismatch on index %v: %v != %v", i, actual[i], expected[i])
		}
	}
}

func TestMaximumSumSubarray(t *testing.T) {
	a := []int{10, -1, 2, 5, 11, 2, 0, -10, 9, 1}
	expected := []int{10, -1, 2, 5, 11, 2}

	assertSlicesEqual(t, MaximumSumSubarray(a), expected)
}

func TestMaximumSumSubarrayEmptyArray(t *testing.T) {
	a := []int{}
	expected := []int{}

	assertSlicesEqual(t, MaximumSumSubarray(a), expected)
}

func TestMaximumSumSubarraySortedAscending(t *testing.T) {
	a := make([]int, 0, 21)
	for i := -10; i <= 10; i++ {
		a = append(a, i)
	}

	expected := make([]int, 0, 10)
	for i := 1; i <= 10; i++ {
		expected = append(expected, i)
	}

	assertSlicesEqual(t, MaximumSumSubarray(a), expected)
}

func TestMaximumSumSubarraySortedDescending(t *testing.T) {
	a := make([]int, 0, 21)
	for i := 10; i >= -10; i-- {
		a = append(a, i)
	}

	expected := make([]int, 0, 10)
	for i := 10; i >= 1; i-- {
		expected = append(expected, i)
	}

	assertSlicesEqual(t, MaximumSumSubarray(a), expected)
}

func TestMaximumSumSubarrayAllDuplicates(t *testing.T) {
	a := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		a = append(a, 7)
	}

	expected := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		expected = append(expected, 7)
	}

	assertSlicesEqual(t, MaximumSumSubarray(a), expected)
}

func TestMaximumSumSubarrayAllZeros(t *testing.T) {
	a := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		a = append(a, 0)
	}

	expected := []int{0}

	assertSlicesEqual(t, MaximumSumSubarray(a), expected)
}

func TestMaximumSumSubarrayAllNegative(t *testing.T) {
	a := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		a = append(a, -7)
	}

	expected := []int{-7}

	assertSlicesEqual(t, MaximumSumSubarray(a), expected)
}
