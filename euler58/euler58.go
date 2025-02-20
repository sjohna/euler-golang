package euler58

import (
	"euler/utilities"
	"fmt"
)

/*
https://projecteuler.net/problem=58

Starting with 1 and spiralling anticlockwise, we put all natural numbers into a spiral that starts like this:

37 36 35 34 33 32 31
38 17 16 15 14 13 30
39 18  5  4  3 12 29
40 19  6  1  2 11 28
41 20  7  8  9 10 27
42 21 22 23 24 25 26
43 44 45 46 47 48 49

There seem to be a lot of primes on the diagonals (except for the bottom-right, which contains the odd squares).

What is the smallest size of this construction where the portion of primes on the diagonals is under 10%?
*/

func Euler58() {
	numsOnDiagonal := 1 // center
	currNum := 1
	layer := 2
	primesOnDiagonal := 0

	for {
		for i := 0; i < 4; i++ {
			currNum += layer
			if utilities.IsPrime(currNum) {
				primesOnDiagonal++
			}
		}
		numsOnDiagonal += 4

		percent := float64(primesOnDiagonal) / float64(numsOnDiagonal)

		fmt.Printf("layer %d: %d/%d (%f)\n", layer, primesOnDiagonal, numsOnDiagonal, percent)

		if percent < .10 {
			break
		}

		layer += 2
	}

	fmt.Println(layer + 1)
}
