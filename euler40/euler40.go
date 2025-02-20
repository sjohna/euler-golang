package euler40

import (
	"strconv"
)

/*
https://projecteuler.net/problem=40

Champernowne's constant is the number whose decimal representation is the concatenation of the positive integers:

0.12345678910111213...

What is the product of the (10^n)th digits of Champernowne's constant, for n in 0..6
*/

func ChampernowneDigit(n int) int {
	lengthSoFar := 0
	numsInNextPowerBlock := 9
	length := 1

	for {
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
}

func Basic() int {
	product := 1
	for i := 1; i <= 1000000; i *= 10 {
		product *= ChampernowneDigit(i)
	}

	return product
}
