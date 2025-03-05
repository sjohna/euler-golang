package euler14

/*
https://projecteuler.net/problem=14

What is the length of the longest Collatz sequence starting from a number under one million?
*/

func RecursiveWithCache() int {
	cache := make(map[int]int)
	cache[1] = 1

	maxLength := 0
	maxValue := -1

	for i := 1; i < 1_000_000; i++ {
		if length(cache, i) > maxLength {
			maxLength = length(cache, i)
			maxValue = i
		}
	}

	return maxValue
}

func length(cache map[int]int, n int) int {
	if length, ok := cache[n]; ok {
		return length

	}

	cache[n] = 1 + length(cache, nextValue(n))

	return cache[n]
}

func nextValue(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
