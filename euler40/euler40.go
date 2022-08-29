package euler40

import (
	"fmt"
	"strconv"
)

func ChampernowneDigit(n int) int {
	lengthSoFar := 0
	numsInNextPowerBlock := 9
	length := 1

	for true {
		lengthOfNextBlock := numsInNextPowerBlock * length
		if n <= lengthSoFar+lengthOfNextBlock {
			offsetIntoBlock := n - lengthSoFar
			numInBlock := offsetIntoBlock / length
			offsetInNum := numInBlock % length

			// calculate power of 10 to add
			min := 1
			for i := 1; i < length; i++ {
				min *= 10
			}

			if min == 1 {
				min = 0
			}

			num := min + numInBlock
			numStr := strconv.Itoa(num)

			digitStr := string(rune(numStr[offsetInNum]))
			val, _ := strconv.Atoi(digitStr)
			return val
		} else {
			lengthSoFar += lengthOfNextBlock
			numsInNextPowerBlock *= 10
			length++
		}
	}

	return -1
}

func main() {
	product := 1
	for i := 1; i <= 1000000; i *= 10 {
		product *= ChampernowneDigit(i)
	}

	fmt.Println(product)
}
