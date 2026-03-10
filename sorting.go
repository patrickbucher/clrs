package clrs

import "cmp"

func InsertionSort[T cmp.Ordered](items []T) {
	n := len(items)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		key := items[i]
		j := i - 1
		for ; j >= 0 && items[j] > key; j-- {
			items[j+1] = items[j]
		}
		items[j+1] = key
	}
}
