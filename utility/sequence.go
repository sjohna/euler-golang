package utility

func NaturalNumbers() Generator[int] {
	curr := 0

	return func() (int, bool) {
		curr++
		return curr, true
	}
}

func FibonacciSequence() Generator[int] {
	curr := 0
	next := 1

	return func() (int, bool) {
		curr, next = next, next+curr
		return curr, true
	}
}

// Range a range of integers, inclusive on both ends
func Range(min, max int) Generator[int] {
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
