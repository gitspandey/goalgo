package array

// Reference: https://en.wikipedia.org/wiki/Binary_search_algorithm

// Custom accessor to allow stubbing for unit tests
type BinarySearcher struct {
	// Return element at given index
	ElementAt func([]int, int) int
}

func (searcher *BinarySearcher) Search(a []int, value int) int {
	elementAt := searcher.ElementAt
	if elementAt == nil {
		elementAt = func(a []int, i int) int {
			return a[i]
		}
	}

	left, right := 0, len(a)-1
	for left <= right {
		mid := (left + right) / 2
		t := elementAt(a, mid)

		if t == value {
			return t
		}

		if t < value {
			left = mid + 1
		} else if t > value {
			right = mid - 1
		}
	}
	return -1
}
