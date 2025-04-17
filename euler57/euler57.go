package euler57

import (
	"fmt"
	"math/big"
)

func Iterate() int {
	one := new(big.Rat).SetInt64(1)
	two := new(big.Rat).SetInt64(2)

	fracPart := new(big.Rat).Quo(one, two)

	numLonger := 0

	for range 1000 {
		thisExpansion := new(big.Rat).Add(one, fracPart)
		fracPart = fracPart.Quo(one, new(big.Rat).Add(two, fracPart))
		numeratorLength := len(thisExpansion.Num().Text(10))
		denominatorLength := len(thisExpansion.Denom().Text(10))
		fmt.Printf("%s", thisExpansion.RatString())
		if numeratorLength > denominatorLength {
			numLonger++
			fmt.Printf(" +")
		}
		fmt.Println()
	}

	return numLonger
}
