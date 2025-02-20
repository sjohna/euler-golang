package euler4

import (
	"euler/utilities"
	"sort"
)

/*
https://projecteuler.net/problem=4

Find the largest palindrome make from the product of two 3-digit numbers.
*/

func Loops() int {
	products := make([]int, 0)

	for n1 := 1; n1 < 1000; n1++ {
		for n2 := n1; n2 < 1000; n2++ {
			products = append(products, n1*n2)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(products)))

	for i := 0; i < len(products); i++ {
		if utilities.IsPalindrome(products[i]) {
			return products[i]
		}
	}

	return -1
}
