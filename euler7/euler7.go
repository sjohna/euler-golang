package euler7

import (
	"euler/utility/prime"
)

/*
https://projecteuler.net/problem=7

What is the 10001st prime number?
*/

func Generator() int {
	return prime.CachedGenerator().Skip(10000).NextValue()
}
