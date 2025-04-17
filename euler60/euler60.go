package euler60

import (
	"euler/utility"
	"slices"
)

// slow: 3-5 minutes
func SolveIt() int {
	cliques := make([][]int, 0)
	primes := make(map[int]bool)
	primeGen := utility.PrimeGenerator().Infinite()
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
				pre := utility.Concat(nextPrime, p)
				post := utility.Concat(p, nextPrime)
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
					return utility.SliceGenerator(newClique).Reduce(utility.Sum, 0)
				}
				newCliques = append(newCliques, newClique)
			}
		}

		cliques = append(cliques, newCliques...)
		currPrime = nextPrime
	}
}
