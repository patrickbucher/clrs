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

func SelectionSort[T cmp.Ordered](items []T) {
	n := len(items)
	if n <= 1 {
		return
	}
	for i := range n - 1 {
		k := i + 1
		smallest := items[k]
		for j := k + 1; j < n; j++ {
			if items[j] < smallest {
				smallest = items[j]
				k = j
			}
		}
		if smallest < items[i] {
			temp := items[i]
			items[i] = items[k]
			items[k] = temp
		}
	}
}

func MergeSort[T cmp.Ordered](items []T) {
	n := len(items)
	if n <= 1 {
		return
	}
	mergeSort(items, 0, n-1)
}

func mergeSort[T cmp.Ordered](items []T, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort(items, p, q)
	mergeSort(items, q+1, r)
	merge(items, p, q, r)
}

func merge[T cmp.Ordered](items []T, p, q, r int) {
	nl := q - p + 1
	nr := r - q
	left := make([]T, nl)
	right := make([]T, nr)
	copy(left, items[p:q+1])
	copy(right, items[q+1:r+1])
	i, j, k := 0, 0, p
	for i < nl && j < nr {
		if left[i] <= right[j] {
			items[k] = left[i]
			i++
		} else {
			items[k] = right[j]
			j++
		}
		k++
	}
	for i < nl {
		items[k] = left[i]
		i++
		k++
	}
	for j < nr {
		items[k] = right[j]
		j++
		k++
	}
}
