package main

import (
	"euler"
	"fmt"
)

func Euler2() int {
	nextFib := euler.FibonacciGenerator()
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
