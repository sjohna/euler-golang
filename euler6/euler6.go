package euler6

import "euler/utility"

/*
https://projecteuler.net/problem=6

What's the difference between the square of the sum and the sum of the squares of the first 100 natural numbers?
*/

func Generators() int {
	sumOfSquares := utility.Range(1, 100).Map(utility.Square).Reduce(utility.Sum, 0)
	squareOfSum := utility.Square(utility.Range(1, 100).Reduce(utility.Sum, 0))

	return squareOfSum - sumOfSquares
}
