package utilities

func TakeN[T comparable](generator func() T, N int) []T {
	ret := make([]T, N)
	for i := 0; i < N; i++ {
		ret[i] = generator()
	}

	return ret
}

func Nth[T any](generator func() T, N int) T {
	for i := 0; i < N-1; i++ {
		generator()
	}

	return generator()
}
