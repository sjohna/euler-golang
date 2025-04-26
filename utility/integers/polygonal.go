package integers

func Triangular(n int) int {
	return (n * (n + 1)) / 2
}

func Pentagonal(n int) int {
	return (n * (3*n - 1)) / 2
}

func Hexagonal(n int) int {
	return n * (2*n - 1)
}

func Heptagonal(n int) int {
	return (n * (5*n - 3)) / 2
}

func Octagonal(n int) int {
	return n * (3*n - 2)
}
