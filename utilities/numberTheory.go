package utilities

func GCD(a int, b int) int {
	if a < b {
		a, b = b, a
	}

	for ; b != 0; a, b = b, a%b {
	}

	return a
}

func LCM(a int, b int) int {
	return (a * b) / GCD(a, b)
}
