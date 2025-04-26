package euler60

import (
	"euler/utility"
	"euler/utility/generator"
	"euler/utility/integers"
	"euler/utility/prime"
	"slices"
)

// slow: 3-5 minutes
func SolveIt() int {
	cliques := make([][]int, 0)
	primes := make(map[int]bool)
	primeGen := prime.NaiveGenerator().Infinite()
	nextPrimeMap := make(map[int]int)
	currPrime := -1
	maxPrime := -1

	for {
		var ok bool
		var nextPrime int
		if nextPrime, ok = nextPrimeMap[currPrime]; !ok {
			nextPrime = primeGen()
			nextPrimeMap[currPrime] = nextPrime
			maxPrime = nextPrime
		}

		newCliques := make([][]int, 0)
		newCliques = append(newCliques, []int{nextPrime})

		for _, c := range cliques {
			allWork := true
			for _, p := range c {
				pre := integers.Concat(nextPrime, p)
				post := integers.Concat(p, nextPrime)
				for maxPrime < pre || maxPrime < post {
					newMax := primeGen()
					nextPrimeMap[maxPrime] = newMax
					primes[newMax] = true
					maxPrime = newMax
				}

				if !primes[pre] || !primes[post] {
					allWork = false
					break
				}
			}

			if allWork {
				newClique := append(slices.Clone(c), nextPrime)
				if len(newClique) == 5 {
					return generator.Slice(newClique).Reduce(utility.Sum, 0)
				}
				newCliques = append(newCliques, newClique)
			}
		}

		cliques = append(cliques, newCliques...)
		currPrime = nextPrime
	}
}
