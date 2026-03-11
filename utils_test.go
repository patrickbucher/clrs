package clrs

import (
	"slices"
	"testing"
)

func TestRandIntsEmpty(t *testing.T) {
	actual := RandInts(9, 1, 10)
	expected := []int{}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected empty slice, was %v\n", actual)
	}
}

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
			t.Fatalf("ints[%d] = %d is out of range [%d,%d]\n",
				i, v, min, max)
		}
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		items  []int
		sorted bool
	}{
		{[]int{}, true},
		{[]int{0}, true},
		{[]int{0, 0}, true},
		{[]int{0, 1}, true},
		{[]int{0, 0}, true},
		{[]int{1, 0}, false},
		{[]int{0, 9, 10}, true},
		{[]int{0, 10, 9}, false},
	}
	for _, test := range tests {
		actual := IsSorted(test.items)
		if actual != test.sorted {
			t.Errorf("expected Sorted(%v) to be %v, was %v\n",
				test.items, test.sorted, actual)
		}
	}
}
