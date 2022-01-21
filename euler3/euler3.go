package main

import (
	"euler"
	"fmt"
)

func Euler3() int {
	curr := 600851475143
	nextPrime := euler.PrimeGenerator()
	var currPrime int

	for curr > 1 {
		currPrime = nextPrime()
		for curr%currPrime == 0 {
			curr /= currPrime
		}
	}

	return currPrime
}

func main() {
	fmt.Println(Euler3())
}
