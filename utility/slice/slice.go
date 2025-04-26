package slice

import (
	"euler/utility/types"
)

func Sum[T types.Number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}

	return sum
}

func Product[T types.Number](slice []T) T {
	var product T = 1
	for _, v := range slice {
		product *= v
	}

	return product
}

func IsPalindrome[T comparable](slice []T) bool {
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] != slice[len(slice)-i-1] {
			return false
		}
	}
	return true
}

func Select[T any](slice []T, indices ...int) []T {
	ret := make([]T, len(indices))
	for retIndex, index := range indices {
		ret[retIndex] = slice[index]
	}

	return ret
}
