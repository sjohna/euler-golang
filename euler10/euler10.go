package euler10

import "euler/utilities"

func Generator() int {
	return utilities.PrimeGenerator().TakeWhile(utilities.LessThan(2_000_000)).Reduce(utilities.Sum, 0)
}
