package sort

// Reference: https://en.wikipedia.org/wiki/Bubble_sort

func BubbleSort (a []int) {
	swapped := true
	for end := len(a); end > 0 && swapped; end-- {
		swapped = false
		for i := 0; i < end - 1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
	}
}
