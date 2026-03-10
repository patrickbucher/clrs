package clrs

import (
	"cmp"
	"math/rand"
)

func RandInts(min, max, n int) []int {
	if min > max {
		return []int{}
	}
	ints := make([]int, n)
	for i := range n {
		ints[i] = rand.Intn(max-min+1) + min
	}
	return ints
}

func IsSorted[T cmp.Ordered](items []T) bool {
	n := len(items)
	if n <= 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if cmp.Compare(items[i-1], items[i]) == 1 {
			return false
		}
	}
	return true
}
