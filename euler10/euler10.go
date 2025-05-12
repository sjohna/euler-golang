package euler10

import (
	"euler/utility"
	"euler/utility/prime"
)

/*
https://projecteuler.net/problem=10

Find the sum of all the primes below two million.
*/

func Generator() int {
	return prime.SievePrimesUpTo(2_000_000).Reduce(utility.Sum, 0)
}
