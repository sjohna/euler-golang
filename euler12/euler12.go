package euler12

import (
	"euler/utility/prime"
	"euler/utility/sequence"
)

/*
https://projecteuler.net/problem=12

Find the first triangular number with 500 or more factors.
*/

func Generator() int {
	return sequence.TriangularNumbers().Filter(func(n int) bool { return prime.TotalFactors(n) >= 500 }).NextValue()
}
