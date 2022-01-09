package main

import "fmt"

func FibonacciSequence() func() int {
	curr := 0
	next := 1

	return func() int {
		curr, next = next, next+curr
		return curr
	}
}

func Euler2() int {
	nextFib := FibonacciSequence()
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

func main() {
	fmt.Println(Euler2())
}
