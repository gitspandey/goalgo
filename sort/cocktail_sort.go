package sort

// Reference: https://en.wikipedia.org/wiki/Cocktail_shaker_sort

func CocktailSort(a []int) {
	swapped := true
	for start := 0; start < len(a) && swapped; start++ {
		swapped = false
		for i := start; i < len(a)-start-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false
		for i := len(a) - start - 2; i > start; i-- {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				swapped = true
			}
		}
	}
}
