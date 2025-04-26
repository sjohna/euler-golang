package prime

import (
	"euler/utility/generator"
	"math"
)

func NaiveGenerator() generator.Generator[int] {
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

func CachedGenerator() generator.Generator[int] {
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

// PrimeFactorization returns a map of the multiplicities of each prime in the factorization
func PrimeFactorization(num int) map[int]int {
	// this regenerates all the primes again each time. Not ideal
	primes := CachedGenerator().Infinite()

	factorization := make(map[int]int)
	for num > 1 {
		nextPrime := primes()

		mult := 0
		for num%nextPrime == 0 {
			mult += 1
			num /= nextPrime
		}

		if mult > 0 {
			factorization[nextPrime] = mult
		}
	}

	return factorization
}

// NextPrime returns the next larger prime, given a list of the first n primes. Assumes list is sorted in increasing order.
func NextPrime(primes []int) int {
	if len(primes) == 0 {
		return 2
	}
	currPrime := primes[len(primes)-1]
	if currPrime == 2 {
		return 3
	}

	nextCandidate := currPrime + 2

	for {
		isPrime := false
		maxFactor := int(math.Sqrt(float64(nextCandidate)))
		for _, prime := range primes {
			if prime > maxFactor {
				break
			}

			if nextCandidate%prime == 0 {
				isPrime = true
				break
			}
		}

		if !isPrime {
			return nextCandidate
		}

		nextCandidate += 2
	}
}

var primeFactorCache []int

func PrimeFactorizationWithCache(num int) map[int]int {
	factorization := make(map[int]int)
	for _, prime := range primeFactorCache {
		mult := 0
		for num%prime == 0 {
			mult += 1
			num /= prime
		}

		if mult > 0 {
			factorization[prime] = mult
		}

		if num == 1 {
			break
		}
	}

	if num == 1 {
		return factorization
	}

	// else, we exhausted the prime list without finding all the factors. Find more primes until we do.
	for {
		nextPrime := NextPrime(primeFactorCache)
		primeFactorCache = append(primeFactorCache, nextPrime)

		mult := 0
		for num%nextPrime == 0 {
			mult += 1
			num /= nextPrime
		}

		if mult > 0 {
			factorization[nextPrime] = mult
		}

		if num == 1 {
			break
		}
	}

	return factorization
}

func TotalFactors(num int) int {
	factorization := PrimeFactorizationWithCache(num)
	total := 1

	for _, mult := range factorization {
		total *= mult + 1
	}

	return total
}

// TODO: test
func MultiplyPrimeFactorizations(a, b map[int]int) map[int]int {
	product := make(map[int]int)
	for prime, mult := range a {
		product[prime] += mult
	}

	for prime, mult := range b {
		product[prime] += mult
	}

	return product
}

func SievePrimesUpTo(n int) generator.Generator[int] {
	sieve := make([]bool, n+1)
	sieve[0] = true
	sieve[1] = true

	startIndex := 0

	return func() (int, bool) {
		// find next prime
		for startIndex < len(sieve) {
			if !sieve[startIndex] {
				break
			}
			startIndex++
		}

		if startIndex >= len(sieve) {
			return 0, false
		}

		nextPrime := startIndex

		for i := startIndex * 2; i < len(sieve); i += startIndex {
			sieve[i] = true
		}

		startIndex++
		return nextPrime, true
	}

}
