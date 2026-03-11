package clrs

import (
	"errors"
	"slices"
	"testing"
)

func TestBinaryAddition(t *testing.T) {
	tests := []struct {
		a []int
		b []int
		c []int
		e error
	}{
		{[]int{}, []int{}, []int{0}, nil},
		{[]int{}, []int{0}, nil, ErrLengthMismatch},
		{[]int{0}, []int{}, nil, ErrLengthMismatch},
		{[]int{0}, []int{0}, []int{0, 0}, nil},
		{[]int{0}, []int{1}, []int{0, 1}, nil},
		{[]int{1}, []int{0}, []int{0, 1}, nil},
		{[]int{1}, []int{1}, []int{1, 0}, nil},
		{[]int{0, 0}, []int{0, 0}, []int{0, 0, 0}, nil},
		{[]int{0, 1}, []int{0, 1}, []int{0, 1, 0}, nil},
		{[]int{1, 0}, []int{0, 1}, []int{0, 1, 1}, nil},
		{[]int{1, 1}, []int{0, 1}, []int{1, 0, 0}, nil},
		{[]int{1, 1}, []int{1, 1}, []int{1, 1, 0}, nil},
		{[]int{0, 0}, []int{0}, nil, ErrLengthMismatch},
		{[]int{0}, []int{0, 0}, nil, ErrLengthMismatch},
		{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0, 0}, nil},
		{[]int{0, 0, 1}, []int{0, 0, 0}, []int{0, 0, 0, 1}, nil},
		{[]int{0, 0, 1}, []int{0, 0, 1}, []int{0, 0, 1, 0}, nil},
		{[]int{0, 1, 0}, []int{0, 0, 1}, []int{0, 0, 1, 1}, nil},
		{[]int{0, 1, 0}, []int{0, 1, 1}, []int{0, 1, 0, 1}, nil},
		{[]int{0, 1, 1}, []int{0, 1, 1}, []int{0, 1, 1, 0}, nil},
		{[]int{0, 1, 1}, []int{1, 1, 1}, []int{1, 0, 1, 0}, nil},
	}
	for _, test := range tests {
		a := reversed(test.a)
		b := reversed(test.b)
		expected := reversed(test.c)
		actual, err := BinaryAddition(a, b)
		if !slices.Equal(expected, actual) {
			t.Errorf("%v + %v: expected %v, was %v\n",
				test.a, test.b, test.c, reversed(actual))
		}
		if err != nil && test.e != nil && !errors.Is(err, test.e) {
			t.Errorf("%v + %v: expected error %v, was %v\n",
				test.a, test.b, test.e, err)
		}
	}
}

func reversed[T any](s []T) []T {
	n := len(s)
	c := make([]T, n)
	copy(c, s)
	slices.Reverse(c)
	return c
}
