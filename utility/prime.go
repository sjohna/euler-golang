package utility

import "math"

func PrimeGenerator() Generator[int] {
	curr := 2

	return func() (int, bool) {
		if curr == 2 {
			curr = 3
			return 2, true
		} else {
			for {
				if curr%2 == 0 {
					curr += 2
					continue
				}

				isPrime := true
				factorMax := int(math.Sqrt(float64(curr)))
				for factor := 3; factor <= factorMax; factor += 2 {
					if curr%factor == 0 {
						isPrime = false
						break
					}
				}

				if isPrime {
					break
				} else {
					curr += 2
				}
			}
		}

		ret := curr
		curr += 2
		return ret, true
	}
}

func CachedPrimeGenerator() Generator[int] {
	curr := 2
	primes := make([]int, 0)

	return func() (int, bool) {
		if curr == 2 {
			curr = 3
			primes = append(primes, 2)
			return 2, true
		} else {
			for {
				isPrime := true
				factorMax := int(math.Sqrt(float64(curr)))
				for _, factor := range primes {
					if factor > factorMax {
						break
					}

					if curr%factor == 0 {
						isPrime = false
						break
					}

				}

				if isPrime {
					break
				} else {
					curr += 2
				}
			}
		}

		ret := curr
		curr += 2
		primes = append(primes, ret)
		return ret, true
	}
}

func IsPrime(num int) bool {
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	factorMax := int(math.Sqrt(float64(num)))
	for factor := 3; factor <= factorMax; factor += 2 {
		if num%factor == 0 {
			return false
		}
	}

	return true
}
