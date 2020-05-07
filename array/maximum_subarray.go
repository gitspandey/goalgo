package array

// Reference: https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm

func MaximumSumSubarray(a []int) []int {
	if len(a) == 0 {
		return a
	}

	// Scan the array from left to right, keeping track of the max sum ending at each index
	// as well as the max sum seen so far

	maxsum, maxstart, maxend := 0, 0, 0
	cursum, curstart := 0, 0
	for curend, n := range a {
		if cursum <= 0 {
			// start a new sum at this number
			curstart = curend
			cursum = 0
		}
		cursum += n

		if cursum > maxsum {
			maxsum, maxstart, maxend = cursum, curstart, curend
		}
	}
	return a[maxstart : maxend+1]
}
