package sort

// Reference: https://en.wikipedia.org/wiki/Gnome_sort

func GnomeSort (a []int) {
	for i := 0; i < len(a); {
		if i == 0 || a[i] >= a[i-1] {
			i++
		} else {
			a[i-i], a[i] = a[i], a[i-1]
			i--
		}
	}
}
