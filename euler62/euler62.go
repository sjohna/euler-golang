package euler62

import (
	"euler/utility/sequence"
	"slices"
)

func SortDigits(n int) int {
	digitCounts := make(map[int]int)

	for n > 0 {
		firstDigit := n % 10
		digitCounts[firstDigit]++
		n = n / 10
	}

	sorted := 0

	for d := 9; d >= 0; d-- {
		for range digitCounts[d] {
			sorted *= 10
			sorted += d
		}
	}

	return sorted
}

func SolveIt() int {
	cube := func(n int) int {
		return n * n * n
	}

	cubeGen := sequence.NaturalNumbers().Map(cube).Infinite()

	sortedMap := make(map[int][]int)

	for {
		currCube := cubeGen()
		currCubeSorted := SortDigits(currCube)

		if cubes, ok := sortedMap[currCubeSorted]; ok {
			sortedMap[currCubeSorted] = append(cubes, currCube)
			if len(sortedMap[currCubeSorted]) == 5 {
				return slices.Min(sortedMap[currCubeSorted])
			}
		} else {
			sortedMap[currCubeSorted] = []int{currCube}
		}
	}
}
