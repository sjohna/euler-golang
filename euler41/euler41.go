package euler41

import (
	"euler/utility"
	"strconv"
	"strings"
)

func Loop() int {
	digitStrings := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	largestPrime := 0

	for n := 9; n > 0; n-- {
		numPermutations := utility.Factorial(n)
		for k := range numPermutations {
			perm := utility.KthPermutationOfNItems(n, k+1)
			num, _ := strconv.Atoi(strings.Join(utility.SliceSelect(digitStrings, perm...), ""))

			if utility.IsPrime(num) && num > largestPrime {
				largestPrime = num
			}
		}

		if largestPrime > 0 {
			break
		}
	}

	return largestPrime
}
