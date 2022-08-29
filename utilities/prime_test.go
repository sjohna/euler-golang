package utilities

import "testing"

func TestPrimeGeneratorFirstValues(t *testing.T) {
	nextPrime := PrimeGenerator()

	if val, expected := nextPrime(), 2; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 3; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 5; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 7; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 11; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}
}

func TestCachedPrimeGeneratorFirstValues(t *testing.T) {
	nextPrime := CachedPrimeGenerator()

	if val, expected := nextPrime(), 2; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 3; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 5; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 7; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextPrime(), 11; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}
}

func TestIsPrime(t *testing.T) {
	if val := 2; !IsPrime(val) {
		t.Fatalf("%d is prime but IsPrime returned false", val)
	}

	if val := 3; !IsPrime(val) {
		t.Fatalf("%d is prime but IsPrime returned false", val)
	}

	if val := 5; !IsPrime(val) {
		t.Fatalf("%d is prime but IsPrime returned false", val)
	}

	if val := 7; !IsPrime(val) {
		t.Fatalf("%d is prime but IsPrime returned false", val)
	}

	if val := 11; !IsPrime(val) {
		t.Fatalf("%d is prime but IsPrime returned false", val)
	}

	if val := 31; !IsPrime(val) {
		t.Fatalf("%d is prime but IsPrime returned false", val)
	}

	if val := 89231; !IsPrime(val) {
		t.Fatalf("%d is prime but IsPrime returned false", val)
	}

	if val := 4; IsPrime(val) {
		t.Fatalf("%d is not prime but IsPrime returned true", val)
	}

	if val := 6; IsPrime(val) {
		t.Fatalf("%d is not prime but IsPrime returned true", val)
	}

	if val := 49; IsPrime(val) {
		t.Fatalf("%d is not prime but IsPrime returned true", val)
	}

	if val := 50; IsPrime(val) {
		t.Fatalf("%d is not prime but IsPrime returned true", val)
	}

	if val := 89400; IsPrime(val) {
		t.Fatalf("%d is not prime but IsPrime returned true", val)
	}
}
