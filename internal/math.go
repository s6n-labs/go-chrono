package internal

import "golang.org/x/exp/constraints"

func abs[T constraints.Integer](a T) T {
	if a >= 0 {
		return a
	}

	return -a
}

func remEuclid[T constraints.Integer](a, b T) T {
	r := a % b
	if r >= 0 {
		return r
	}

	return r + abs(b)
}
