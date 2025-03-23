package euler78

import (
	"fmt"
)

func DynamicProgramming() int {
	divisions := make([][]int, 0)
	divisions = append(divisions, []int{1}) // one way to partition no coins

	coinCount := 1
	for {
		nextCounts := make([]int, coinCount+1)
		for i := 1; i <= coinCount; i++ {
			lookBack := coinCount - i
			index := i
			if index >= len(divisions[lookBack]) {
				index = len(divisions[lookBack]) - 1
			}

			nextCounts[i] += divisions[lookBack][index]
			nextCounts[i] += nextCounts[i-1]
			nextCounts[i] %= 1_000_000
		}

		if nextCounts[coinCount] == 0 {
			return coinCount
		}

		if coinCount%10_000 == 0 {
			fmt.Printf("%d: %d\n", coinCount, nextCounts[coinCount])
		}

		divisions = append(divisions, nextCounts)
		coinCount++
	}
}

// EulersPentagonalNumberTheorem based on https://en.wikipedia.org/wiki/Partition_function_(number_theory) and https://en.wikipedia.org/wiki/Pentagonal_number_theorem
func EulersPentagonalNumberTheorem() int {
	partitions := []int{1} // 1 way to partition 0 coins
	n := 1
	for {
		partitionsOfN := 0
		k := 1
		sign := 1
		for {
			diff1 := (k * (3*k - 1)) / 2
			diff2 := (k * (3*k + 1)) / 2

			used := false

			if n-diff1 >= 0 {
				partitionsOfN += sign * partitions[n-diff1]
				partitionsOfN += 1_000_000
				partitionsOfN %= 1_000_000
				used = true
			}

			if n-diff2 >= 0 {
				partitionsOfN += sign * partitions[n-diff2]
				partitionsOfN += 1_000_000
				partitionsOfN %= 1_000_000
				used = true
			}

			if !used {
				break
			}

			k++
			sign *= -1
		}

		if partitionsOfN%1_000_000 == 0 {
			return n
		}

		partitions = append(partitions, partitionsOfN)
		n++
	}
}
