package euler

import (
	"strconv"
)

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

func Reverse(num int) int {
	numstr := strconv.Itoa(num)
	runes := []rune(numstr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	val, _ := strconv.Atoi(string(runes))
	return val
}
