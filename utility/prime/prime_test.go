package prime

import (
	"euler/utility/generator"
	"reflect"
	"testing"
)

func testPrimeGeneratorFirstValues(t *testing.T, gen generator.Generator[int]) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	got := gen.Take(10).ToSlice()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("First Values: expected %v, got %v", expected, got)
	}
}

func testPrimeGenerator100th(t *testing.T, gen generator.Generator[int]) {
	expected := 541
	got := gen.Nth(100)

	if expected != got {
		t.Errorf("100th Value: expected %v, got %v", expected, got)
	}
}

func testPrimeGenerator1000th(t *testing.T, gen generator.Generator[int]) {
	expected := 7919
	got := gen.Nth(1000)

	if expected != got {
		t.Errorf("1000th Value: expected %v, got %v", expected, got)
	}
}

func TestPrimeGenerator(t *testing.T) {
	testPrimeGeneratorFirstValues(t, NaiveGenerator())
	testPrimeGenerator100th(t, NaiveGenerator())
	testPrimeGenerator1000th(t, NaiveGenerator())
}

func TestCachedPrimeGenerator(t *testing.T) {
	testPrimeGeneratorFirstValues(t, CachedGenerator())
	testPrimeGenerator100th(t, CachedGenerator())
	testPrimeGenerator1000th(t, CachedGenerator())
}

func TestIsPrime(t *testing.T) {
	type test struct {
		n        int
		expected bool
	}

	tests := []test{
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{11, true},
		{31, true},
		{89231, true},
		{4, false},
		{6, false},
		{49, false},
		{50, false},
		{89400, false},
	}

	for _, test := range tests {
		if got := IsPrime(test.n); got != test.expected {
			t.Errorf("IsPrime(%d): expected %v, got %v", test.n, test.expected, got)
		}
	}
}

func TestPrimeFactorization(t *testing.T) {
	type test struct {
		n        int
		expected map[int]int
	}

	tests := []test{
		{1, map[int]int{}},
		{2, map[int]int{2: 1}},
		{3, map[int]int{3: 1}},
		{4, map[int]int{2: 2}},
		{5, map[int]int{5: 1}},
		{6, map[int]int{2: 1, 3: 1}},
		{7, map[int]int{7: 1}},
		{8, map[int]int{2: 3}},
		{9, map[int]int{3: 2}},
		{10, map[int]int{2: 1, 5: 1}},
		{31, map[int]int{31: 1}},
		{89231, map[int]int{89231: 1}},
		{2 * 2 * 3 * 3 * 5 * 5, map[int]int{2: 2, 3: 2, 5: 2}},
		{17 * 31 * 47 * 47, map[int]int{17: 1, 31: 1, 47: 2}},
	}

	type implementation struct {
		name     string
		function func(int) map[int]int
	}

	implementations := []implementation{
		{
			"Naive",
			PrimeFactorization,
		},
		{
			"WithCache",
			PrimeFactorizationWithCache,
		},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, test := range tests {
				if got := PrimeFactorization(test.n); !reflect.DeepEqual(got, test.expected) {
					t.Errorf("PrimeFactorization(%d): expected %v, got %v", test.n, test.expected, got)
				}
			}
		})
	}
}

func TestNextPrime(t *testing.T) {
	// test that first few primes are as expected
	var primes []int

	for range 10 {
		primes = append(primes, NextPrime(primes))
	}

	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	if !reflect.DeepEqual(expected, primes) {
		t.Errorf("first 10: expected %v, got %v", expected, primes)
	}

	// test that certain values are as expected
	for range 990 { // up to 1000 total primes
		primes = append(primes, NextPrime(primes))
	}

	if primes[99] != 541 {
		t.Errorf("100th: expected 541, got %v", primes[99])
	}

	if primes[999] != 7919 {
		t.Errorf("1000th: expected 7919, got %v", primes[999])
	}
}
