package utility

func GCD[T Integer](a T, b T) T {
	if a < b {
		a, b = b, a
	}

	for ; b != 0; a, b = b, a%b {
	}

	return a
}

func LCM[T Integer](a T, b T) T {
	return (a * b) / GCD(a, b)
}
