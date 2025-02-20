package euler55

import (
	"euler/utilities"
)

/*
https://projecteuler.net/problem=55

Consider the process of adding a number to its reverse, like so:

47 + 74 = 121

We can repeat this process. With some numbers, we get a palindrome very quickly (like 47 after one iteration).
With some numbers, we don't seem to ever get a palindrome, no matter how many iterations of this process we do.

We'll consider a number a Lychrel number if, starting from that number, this process does not yield a palindrome within
50 iterations.

How many numbers under 10,000 are Lychrel numbers?
*/

func IsLychrel(n int) bool {
	currVal := n

	for i := 0; i < 50; i++ {
		currVal = currVal + utilities.Reverse(currVal)
		if utilities.IsPalindrome(currVal) {
			return false
		}
	}

	return true
}

func Loop() int {
	numLychrel := 0

	for i := 0; i < 10000; i++ {
		if IsLychrel(i) {
			numLychrel++
		}
	}

	return numLychrel
}

func Generators() int {
	return utilities.NaturalNumbers().
		TakeWhile(func(n int) bool { return n < 10000 }).
		Filter(IsLychrel).
		Count()
}
