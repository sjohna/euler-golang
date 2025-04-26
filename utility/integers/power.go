package integers

import "euler/utility/types"

func Pow[T types.Integer](base, exp T) T {
	pow := base
	var currExp T = 1
	powers := []T{pow}
	for {
		pow *= pow
		currExp *= 2
		if currExp > exp {
			break
		}
		powers = append(powers, pow)
	}

	pow = 1
	currExp /= 2
	index := len(powers) - 1
	for exp > 0 {
		if currExp <= exp {
			pow *= powers[index]
			exp -= currExp
		}
		currExp /= 2
		index--
	}

	return pow
}

func PowMod[T types.Integer](base, exp, mod T) T {
	pow := base
	pow %= mod
	var currExp T = 1
	powers := []T{pow}
	for {
		pow *= pow
		pow %= mod
		currExp *= 2
		if currExp > exp {
			break
		}
		powers = append(powers, pow)
	}

	pow = 1
	currExp /= 2
	index := len(powers) - 1
	for exp > 0 {
		if currExp <= exp {
			pow *= powers[index]
			pow %= mod
			exp -= currExp
		}
		currExp /= 2
		index--
	}

	return pow
}
