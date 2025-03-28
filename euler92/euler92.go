package euler92

func Recursively() int {
	count := 0
	cache := make([]int, 10_000_001)

	for i := 1; i <= 10_000_000; i++ {
		if OrbitFor(i, cache) == 89 {
			count += 1
		}
	}

	return count
}

func OrbitFor(n int, cache []int) int {
	if cache[n] != 0 {
		return cache[n]
	}

	next := NextVal(n)
	if next == 1 {
		cache[n] = 1
		return next
	}

	if next == 89 {
		cache[n] = 89
		return next
	}

	cache[n] = OrbitFor(NextVal(n), cache)
	return cache[n]
}

func NextVal(n int) int {
	sum := 0

	for n > 0 {
		nextDigit := n % 10
		sum += nextDigit * nextDigit
		n /= 10
	}

	return sum
}
