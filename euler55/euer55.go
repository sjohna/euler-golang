package main

import (
	"euler"
	"fmt"
)

func IsLychrel(n int) bool {
	currVal := n

	for i := 0; i < 50; i++ {
		currVal = currVal + euler.Reverse(currVal)
		if euler.IsPalindrome(currVal) {
			return false
		}
	}

	return true
}

func main() {
	numLychrel := 0

	for i := 0; i < 10000; i++ {
		if IsLychrel(i) {
			numLychrel++
		}
	}

	fmt.Println(numLychrel)
}
