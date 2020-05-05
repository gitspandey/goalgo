package sort

// Reference: https://en.wikipedia.org/wiki/Merge_sort

func MergeSort(a []int) {
	if len(a) < 2 {
		return
	}

	// sort the left and right halves
	mid := len(a) / 2
	MergeSort(a[:mid])
	MergeSort(a[mid:])

	// then, merge the two and copy back into the original list
	dst := make([]int, 0, len(a))
	for left, right := 0, mid; len(dst) < len(a); {
		if left < mid && (right >= len(a) || a[left] < a[right]) {
			dst = append(dst, a[left])
			left++
		} else {
			dst = append(dst, a[right])
			right++
		}
	}
	copy(a, dst)
}
