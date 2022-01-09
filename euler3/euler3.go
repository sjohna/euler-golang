package main

import "fmt"

func Primes() <-chan int {
	ret := make(chan int)

	go func() {
		ret <- 2
		curr := 3
		for true {
			anyDivisors := false
			for divisor := 2; divisor < curr/2; divisor++ {
				if curr%divisor == 0 {
					anyDivisors = true
					break
				}
			}
			if !anyDivisors {
				ret <- curr
			}
			curr += 2
		}
	}()

	return ret
}

func Euler3() int {
	curr := 600851475143
	primes := Primes()
	var currPrime int

	for curr > 1 {
		currPrime = <-primes
		for curr%currPrime == 0 {
			curr /= currPrime
		}
	}

	return currPrime
}

func main() {
	fmt.Println(Euler3())
}
