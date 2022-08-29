package euler58

import (
	"euler/utilities"
	"fmt"
)

func main() {
	numsOnDiagonal := 1 // center
	currNum := 1
	layer := 2
	primesOnDiagonal := 0

	for true {
		for i := 0; i < 4; i++ {
			currNum += layer
			if utilities.IsPrime(currNum) {
				primesOnDiagonal++
			}
		}
		numsOnDiagonal += 4

		percent := float64(primesOnDiagonal) / float64(numsOnDiagonal)

		fmt.Printf("layer %d: %d/%d (%f)\n", layer, primesOnDiagonal, numsOnDiagonal, percent)

		if percent < .10 {
			break
		}

		layer += 2
	}

	fmt.Println(layer + 1)
}
