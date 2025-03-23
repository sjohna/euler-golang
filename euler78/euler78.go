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
