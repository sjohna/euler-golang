package euler

func FibonacciGenerator() func() int {
	curr := 0
	next := 1

	return func() int {
		curr, next = next, next+curr
		return curr
	}
}
