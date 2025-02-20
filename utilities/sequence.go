package utilities

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
