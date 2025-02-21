package euler7

import "euler/utility"

/*
https://projecteuler.net/problem=7

What is the 10001st prime number?
*/

func Generator() int {
	return utility.CachedPrimeGenerator().Skip(10000).NextValue()
}
