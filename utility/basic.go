package utility

import (
	"cmp"
	"euler/utility/types"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Sum[T types.Number](a, b T) T {
	return a + b
}
func Product[T types.Number](a, b T) T {
	return a * b
}
func Square[T types.Number](x T) T {
	return x * x
}

func LessThan[T cmp.Ordered](N T) func(T) bool {
	return func(n T) bool {
		return n < N
	}
}
