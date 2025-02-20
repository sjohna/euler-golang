package utilities

// Generator returns a stream of values, and whether the value is valid
type Generator[T any] func() (T, bool)
type Reduction[T any] func(T, T) T

func TakeN[T comparable](g Generator[T], N int) []T {
	ret := make([]T, N)
	for i := 0; i < N; i++ {
		next, _ := g()
		ret[i] = next
	}

	return ret
}

func Nth[T any](g Generator[T], N int) T {
	for i := 0; i < N-1; i++ {
		g()
	}

	nth, _ := g()
	return nth
}

func Default[T any]() T {
	var ret T
	return ret
}

func (g Generator[T]) Reduce(reduction Reduction[T], initial T) T {
	val := initial
	for next, ok := g(); ok; next, ok = g() {
		val = reduction(next, val)
	}

	return val
}

func (g Generator[T]) TakeWhile(cond func(T) bool) Generator[T] {
	return func() (T, bool) {
		next, ok := g()
		if !ok {
			return Default[T](), false
		}

		if cond(next) {
			return next, true
		} else {
			return Default[T](), false
		}
	}
}

func (g Generator[T]) Filter(cond func(T) bool) Generator[T] {
	return func() (T, bool) {
		for {
			next, ok := g()
			if !ok {
				return Default[T](), false
			}

			if cond(next) {
				return next, true
			}
		}
	}
}

// Infinite allows a generator to be treated as an infinite stream of values, not requiring a check for value available
func (g Generator[T]) Infinite() func() T {
	return func() T {
		next, _ := g()
		return next
	}
}

func Sum[T Number](a, b T) T {
	return a + b
}
