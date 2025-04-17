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

func Factorial(n int) int {
	product := 1
	for i := range n {
		product *= i + 1
	}

	return product
}

func SliceSelect[T any](slice []T, indices ...int) []T {
	ret := make([]T, len(indices))
	for retIndex, index := range indices {
		ret[retIndex] = slice[index]
	}

	return ret
}

func Concat(a, b int) int {
	magnitude := 10
	for magnitude < b {
		magnitude *= 10
	}

	return a*magnitude + b
}
