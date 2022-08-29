package euler3

import (
	"euler/utilities"
)

func Euler3() int {
	curr := 600851475143
	nextPrime := utilities.PrimeGenerator()
	var currPrime int

	for curr > 1 {
		currPrime = nextPrime()
		for curr%currPrime == 0 {
			curr /= currPrime
		}
	}

	return currPrime
}
