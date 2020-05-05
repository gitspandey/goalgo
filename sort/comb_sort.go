package sort

import "math"

// Reference: https://en.wikipedia.org/wiki/Comb_sort

func CombSort(a []int) {
	gap := len(a)
	shrink := 1.3
	sorted := false

	for !sorted {
		gap = int(math.Floor(float64(gap) / shrink))
		if gap <= 1 {
			gap = 1
			sorted = true
		}

		for i := 0; i < len(a) - gap; i++ {
			if a[i] > a[i+gap] {
				a[i], a[i+gap] = a[i+gap], a[i]
				sorted = false
			}
		}
	}
}
