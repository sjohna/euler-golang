package integers

import (
	"euler/utility/types"
	"strconv"
)

func Factorial(n int) int {
	product := 1
	for i := range n {
		product *= i + 1
	}

	return product
}

func GCD[T types.Integer](a T, b T) T {
	if a < b {
		a, b = b, a
	}

	for ; b != 0; a, b = b, a%b {
	}

	return a
}

func LCM[T types.Integer](a T, b T) T {
	return (a * b) / GCD(a, b)
}

func Concat(a, b int) int {
	magnitude := 10
	for magnitude < b {
		magnitude *= 10
	}

	return a*magnitude + b
}

func Reverse(num int) int {
	reverse := 0
	for num > 0 {
		reverse *= 10
		reverse += num % 10
		num /= 10
	}
	return reverse
}

func IsPalindrome(num int) bool {
	numstr := strconv.Itoa(num)
	runes := []rune(numstr) // can I access the characters directly?

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
