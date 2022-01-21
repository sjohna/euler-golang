package euler

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
