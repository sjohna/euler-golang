package euler2

import (
	"euler/utility"
)

/*
https://projecteuler.net/problem=2

Find the sum of the even-valued terms of the Fibonacci sequence below four million.
*/

func GeneratorWithLoop() int {
	nextFib := utility.FibonacciSequence().Infinite()
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

func PureGenerator() int {
	return utility.FibonacciSequence().
		TakeWhile(utility.LessThan(4_000_000)).
		Filter(func(n int) bool { return n%2 == 0 }).
		Reduce(utility.Sum, 0)
}
