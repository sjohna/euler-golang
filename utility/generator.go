package utility

// Generator returns a stream of values, and whether the value is valid.
// Once a generator becomes invalid (i.e. returns false for its second return value), it should never become valid again.
type Generator[T any] func() (T, bool)
type Reduction[T any] func(T, T) T

func Empty[T any]() Generator[T] {
	return func() (T, bool) {
		return Default[T](), false
	}
}

func (g Generator[T]) Take(N int) Generator[T] {
	curr := 0

	return func() (T, bool) {
		if curr >= N {
			return Default[T](), false
		}

		curr++
		return g()
	}
}

func (g Generator[T]) ToSlice() []T {
	ret := make([]T, 0)
	for next, ok := g(); ok; next, ok = g() {
		ret = append(ret, next)
	}

	return ret
}

func (g Generator[T]) Nth(N int) T {
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

func (g Generator[T]) SkipWhile(cond func(T) bool) Generator[T] {
	doneSkipping := false
	return func() (T, bool) {
		next, ok := g()

		for !doneSkipping {
			if !ok {
				return Default[T](), false
			}

			if !cond(next) {
				doneSkipping = true
			} else {
				next, ok = g()
			}
		}

		return next, ok
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

func SliceGenerator[T any](slice []T) Generator[T] {
	return SliceRangeGenerator(slice, 0, len(slice)-1)
}

func PartialSums[T Number](g Generator[T]) Generator[T] {
	sum := Default[T]()
	return func() (T, bool) {
		next, ok := g()
		if !ok {
			return Default[T](), false
		}

		sum += next
		return sum, true
	}
}

// Infinite allows a generator to be treated as an infinite stream of values, not requiring a check for value available
func (g Generator[T]) Infinite() func() T {
	return func() T {
		next, _ := g()
		return next
	}
}
