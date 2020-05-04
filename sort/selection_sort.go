package sort

// Reference: https://en.wikipedia.org/wiki/Selection_sort

// Return the index of the smallest element in a, starting at given index
func minimum (a []int, start int) int {
	m := start
	for i := start + 1; i < len(a); i++ {
		if a[i] < a[m] {
			m = i
		}
	}
	return m
}

func SelectionSort (a []int) {
	for i := 0; i < len(a) - 1; i++ {
		if j := minimum(a, i); j != i {
			a[i], a[j] = a[j], a[i]
		}
	}
}
