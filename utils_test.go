package clrs

import "testing"

func TestRandInts(t *testing.T) {
	min := 11
	max := 20
	n := 10
	ints := RandInts(min, max, n)
	m := len(ints)
	if m != n {
		t.Fatalf("expected length to be %d, was %d\n", n, m)
	}
	for i, v := range ints {
		if v < min || v > max {
			t.Fatalf("ints[%d] = %d is out of range [%d,%d]\n", i, v, min, max)
		}
	}
}
