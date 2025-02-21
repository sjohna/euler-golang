package utility

import "cmp"

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Sum[T Number](a, b T) T {
	return a + b
}
func Product[T Number](a, b T) T {
	return a * b
}
func Square[T Number](x T) T {
	return x * x
}

func LessThan[T cmp.Ordered](N T) func(T) bool {
	return func(n T) bool {
		return n < N
	}
}
