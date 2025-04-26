package euler41

import (
	"euler/utility/combinatorics"
	"euler/utility/integers"
	"euler/utility/prime"
	"euler/utility/slice"
	"strconv"
	"strings"
)

func Loop() int {
	digitStrings := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	largestPrime := 0

	for n := 9; n > 0; n-- {
		numPermutations := integers.Factorial(n)
		for k := range numPermutations {
			perm := combinatorics.KthPermutationOfNItems(n, k+1)
			num, _ := strconv.Atoi(strings.Join(slice.Select(digitStrings, perm...), ""))

			if prime.IsPrime(num) && num > largestPrime {
				largestPrime = num
			}
		}

		if largestPrime > 0 {
			break
		}
	}

	return largestPrime
}
