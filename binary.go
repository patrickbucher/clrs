package clrs

import "errors"

var ErrLengthMismatch = errors.New("length mismatch")

func BinaryAddition(a, b []int) ([]int, error) {
	m := len(a)
	n := len(b)
	if m != n {
		return nil, ErrLengthMismatch
	}
	if n == 0 {
		return []int{0}, nil
	}
	c := make([]int, n+1)
	var carry, i int
	for i = range n {
		bit := a[i] + b[i] + carry
		c[i] = bit % 2
		if bit > 1 {
			carry = 1
		} else {
			carry = 0
		}
	}
	c[i+1] = carry
	return c, nil
}
