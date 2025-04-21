package euler63

import "math"

func SolveIt() int {
	numLogs := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range numLogs {
		numLogs[i] = math.Log10(v)
	}

	count := 0

	min := 0.0
	max := 1.0

	anyInRange := true

	for anyInRange {
		anyInRange = false
		for i := range numLogs {
			if numLogs[i] >= min && numLogs[i] < max {
				anyInRange = true
				count++
			}
			numLogs[i] += math.Log10(float64(i + 1))
		}
		min = max
		max += 1
	}

	return count
}
