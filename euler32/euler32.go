package euler32

import (
	"euler/utility"
	"strconv"
	"strings"
)

/*
https://projecteuler.net/problem=32

This product uses all of the digits 1 through 9 exactly once:

39 x 186 = 7254

What is the sum of all products whose multiplicand/multiplier/product identity is 1-9 pandigital?
*/

func Loop() int {
	products := make(map[int]bool)

	numPermutations := utility.Factorial(9)

	digitStrings := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := range numPermutations {
		permutation := utility.KthPermutationOfNItems(9, i+1)

		// check each place we can put the 'x' and '='

		for xIndex := range 9 {
			for equalsIndex := xIndex + 1; equalsIndex < 9; equalsIndex++ {
				multiplicandSlice := permutation[0 : xIndex+1]
				multiplierSlice := permutation[xIndex+1 : equalsIndex+1]
				productSlice := permutation[equalsIndex+1:]

				if len(multiplicandSlice) == 0 || len(multiplierSlice) == 0 || len(productSlice) == 0 {
					continue
				}

				if len(productSlice) < max(len(multiplicandSlice), len(multiplierSlice)) { // the product is at least as long as its longest factor. Results in a roughly 2x speedup
					continue
				}

				multiplicand, _ := strconv.Atoi(strings.Join(utility.SliceSelect(digitStrings, multiplicandSlice...), ""))
				multiplier, _ := strconv.Atoi(strings.Join(utility.SliceSelect(digitStrings, multiplierSlice...), ""))
				product, _ := strconv.Atoi(strings.Join(utility.SliceSelect(digitStrings, productSlice...), ""))

				if multiplicand*multiplier == product {
					products[product] = true
				}
			}
		}
	}

	sum := 0
	for k, _ := range products {
		sum += k
	}

	return sum
}
