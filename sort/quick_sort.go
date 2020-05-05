package sort

// Reference: https://en.wikipedia.org/wiki/Quicksort

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	pivot := partition(a)
	QuickSort(a[:pivot+1])
	QuickSort(a[pivot+1:])
}

func partition(a []int) int {
	pivot := a[(len(a) - 1) / 2]
	i, j := -1, len(a)
	for {
		for i++; a[i] < pivot; i++{
		}
		for j--; a[j] > pivot; j--{
		}

		if i >= j {
			return j
		}
		a[i], a[j] = a[j], a[i]
	}
}
