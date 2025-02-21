package utility

import "cmp"

// Generator returns a stream of values, and whether the value is valid.
// Once a generator becomes invalid (i.e. returns false for its second return value), it should never become valid again.
type Generator[T any] func() (T, bool)
type Reduction[T any] func(T, T) T

func Empty[T any]() Generator[T] {
	return func() (T, bool) {
		return Default[T](), false
	}
}

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

func (g Generator[T]) Count() int {
	count := 0
	for _, ok := g(); ok; _, ok = g() {
		count++
	}

	return count
}

// Map needs to go to the same type, because Go's type system can't really handle switching the type generically here...
func (g Generator[T]) Map(mapFunc func(T) T) Generator[T] {
	return func() (T, bool) {
		next, ok := g()
		if !ok {
			return Default[T](), false
		}

		return mapFunc(next), ok
	}
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

func (g Generator[T]) Skip(N int) Generator[T] {
	for i := 0; i < N; i++ {
		_, ok := g()
		if !ok {
			return Empty[T]()
		}
	}

	return g
}

func (g Generator[T]) NextValue() T {
	next, _ := g()
	return next
}

func SliceRangeGenerator[T any](slice []T, min, max int) Generator[T] {
	curr := min

	return func() (T, bool) {
		if curr > max {
			return Default[T](), false
		}

		ret := slice[curr]
		curr++
		return ret, true
	}
}

// Infinite allows a generator to be treated as an infinite stream of values, not requiring a check for value available
func (g Generator[T]) Infinite() func() T {
	return func() T {
		next, _ := g()
		return next
	}
}

func LessThan[T cmp.Ordered](N T) func(T) bool {
	return func(n T) bool {
		return n < N
	}
}

func Sum[T Number](a, b T) T {
	return a + b
}
func Product[T Number](a, b T) T {
	return a * b
}
func Square[T Number](x T) T {
	return x * x
}
