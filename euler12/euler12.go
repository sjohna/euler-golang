package euler12

import "euler/utility"

/*
https://projecteuler.net/problem=12

Find the first triangular number with 500 or more factors.
*/

func Generator() int {
	return utility.TriangularNumbers().Filter(func(n int) bool { return utility.TotalFactors(n) >= 500 }).NextValue()
}
