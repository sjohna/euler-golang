package utilities

import (
	"reflect"
	"testing"
)

func testPrimeGeneratorFirstValues(t *testing.T, gen func() int) {
	expected := []int{2, 3, 5, 7, 11, 13}
	got := TakeN(gen, 6)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("First Values: expected %v, got %v", expected, got)
	}
}

func testPrimeGenerator100th(t *testing.T, gen func() int) {
	expected := 541
	got := Nth(gen, 100)

	if expected != got {
		t.Errorf("100th Value: expected %v, got %v", expected, got)
	}
}

func testPrimeGenerator1000th(t *testing.T, gen func() int) {
	expected := 7919
	got := Nth(gen, 1000)

	if expected != got {
		t.Errorf("1000th Value: expected %v, got %v", expected, got)
	}
}

func TestPrimeGenerator(t *testing.T) {
	testPrimeGeneratorFirstValues(t, PrimeGenerator())
	testPrimeGenerator100th(t, PrimeGenerator())
	testPrimeGenerator1000th(t, PrimeGenerator())
}

func TestCachedPrimeGenerator(t *testing.T) {
	testPrimeGeneratorFirstValues(t, CachedPrimeGenerator())
	testPrimeGenerator100th(t, CachedPrimeGenerator())
	testPrimeGenerator1000th(t, CachedPrimeGenerator())
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
