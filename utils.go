package clrs

import "math/rand"

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
