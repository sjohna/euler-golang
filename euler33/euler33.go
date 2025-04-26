package euler33

import (
	"euler/utility/integers"
)

func Loop() int {
	numeratorProduct := 1
	denominatorProduct := 1

	for denominator := 10; denominator < 100; denominator++ {
		if denominator%11 == 0 { // multiples of 11 will be trivial solutions
			continue
		}

		for numerator := 10; numerator < denominator; numerator++ {
			if numerator%11 == 0 {
				continue
			}

			// probably can get away with floating point precision here

			value := float64(numerator) / float64(denominator)

			// two variants to check for ab / cd: a/d and b/c

			if numerator%10 == denominator/10 {
				newNumerator := numerator / 10
				newDenominator := denominator % 10

				if newDenominator != 0 {
					newValue := float64(newNumerator) / float64(newDenominator)
					if value == newValue {
						numeratorProduct *= numerator
						denominatorProduct *= denominator
					}
				}
			}

			if numerator/10 == denominator%10 {
				newNumerator := numerator % 10
				newDenominator := denominator / 10

				if newDenominator != 0 {
					newValue := float64(newNumerator) / float64(newDenominator)
					if value == newValue {
						numeratorProduct *= numerator
						denominatorProduct *= denominator
					}
				}
			}

		}
	}

	return denominatorProduct / integers.GCD(numeratorProduct, denominatorProduct)
}
