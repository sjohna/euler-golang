package euler6

import "euler/utilities"

/*
https://projecteuler.net/problem=6

What's the difference between the square of the sum and the sum of the squares of the first 100 natural numbers?
*/

func Generators() int {
	sumOfSquares := utilities.NaturalNumbers().TakeWhile(func(n int) bool { return n <= 100 }).Map(utilities.Square).Reduce(utilities.Sum, 0)
	squareOfSum := utilities.Square(utilities.NaturalNumbers().TakeWhile(func(n int) bool { return n <= 100 }).Reduce(utilities.Sum, 0))

	return squareOfSum - sumOfSquares
}
