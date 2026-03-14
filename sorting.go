package clrs

import "cmp"
import "fmt"

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
	mergeSort(items, 0, n) // [4, 1, 2, 3], 0, 3
}

func mergeSort[T cmp.Ordered](items []T, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort(items, p, q) // 0, 1
	mergeSort(items, q, r) // 1, 3
	merge(items, p, q, r) // 0, 1, 3
}

func merge[T cmp.Ordered](items []T, p, q, r int) {
	fmt.Printf("merge(%v, %d, %d, %d)\n", items, p, q, r)
	nl := q - p // 1
	nr := r - q // 2
	left := make([]T, nl)
	right := make([]T, nr)
	nc := copy(left, items[p:q+1]) // 0:1
	if nc != nl {
		fmt.Println("left nc expected", nl, "was", nc)
	}
	nc = copy(right, items[q+1:r+1]) // 1:3
	if nc != nr {
		fmt.Println("right nc expected", nr, "was", nc)
	}
	fmt.Println("original", items)
	fmt.Println("left", left)
	fmt.Println("right", right)
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
