package sort

// Reference: https://en.wikipedia.org/wiki/Insertion_sort

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		// keep swapping adjacent elements until a[i] is in the right place
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}
