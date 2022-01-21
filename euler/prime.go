package euler

func PrimeGenerator() func() int {
	curr := 2

	return func() int {
		if curr == 2 {
			curr = 3
			return 2
		} else {
			for true {
				if curr%2 == 0 {
					curr += 2
					continue
				}

				isPrime := true
				for factor := 3; factor < curr/2; factor += 2 {
					if curr%factor == 0 {
						isPrime = false
						break
					}
				}

				if isPrime {
					break
				} else {
					curr += 2
				}
			}
		}

		ret := curr
		curr += 2
		return ret
	}
}
