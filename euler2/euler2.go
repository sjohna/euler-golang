package euler2

import (
	"euler/utilities"
)

func Euler2() int {
	nextFib := utilities.FibonacciGenerator()
	currFib := nextFib()
	sum := 0
	for currFib < 4_000_000 {
		if currFib%2 == 0 {
			sum += currFib
		}
		currFib = nextFib()
	}

	return sum
}
