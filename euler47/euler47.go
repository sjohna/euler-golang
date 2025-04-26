package euler47

import (
	"euler/utility/prime"
)

/*
https://projecteuler.net/problem=47

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?
*/

func Loop() int {
	curr := 2
	streak := 0
	for {
		factorization := prime.PrimeFactorizationWithCache(curr)
		if len(factorization) == 4 {
			streak += 1
		} else {
			streak = 0
		}

		if streak == 4 {
			return curr - 3
		}

		curr += 1
	}
}
