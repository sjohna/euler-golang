package sequence

import (
	"euler/utility/generator"
	"euler/utility/slice"
)

func NaturalNumbers() generator.Generator[int] {
	curr := 0

	return func() (int, bool) {
		curr++
		return curr, true
	}
}

func Fibonacci() generator.Generator[int] {
	curr := 0
	next := 1

	return func() (int, bool) {
		curr, next = next, next+curr
		return curr, true
	}
}

func GeneralizedFibonacci(values ...int) generator.Generator[int] {
	nextIndex := 0
	finishedWithInitialValues := false

	return func() (int, bool) {
		if !finishedWithInitialValues {
			next := values[nextIndex]
			nextIndex++
			if nextIndex >= len(values) {
				finishedWithInitialValues = true
				nextIndex = 0
			}
			return next, true
		}
		next := slice.Sum(values)
		values[nextIndex] = next
		nextIndex = (nextIndex + 1) % len(values)
		return next, true
	}
}

// Range a range of integers, inclusive on both ends
func Range(min, max int) generator.Generator[int] {
	curr := min

	return func() (int, bool) {
		if curr > max {
			return 0, false
		}

		ret := curr
		curr++

		return ret, true
	}
}

func TriangularNumbers() generator.Generator[int] {
	return generator.PartialSums(NaturalNumbers())
}
