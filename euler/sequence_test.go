package euler

import "testing"

func TestFibonacciGeneratorFirstValues(t *testing.T) {
	nextFib := FibonacciGenerator()

	if val, expected := nextFib(), 1; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextFib(), 1; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextFib(), 2; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextFib(), 3; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextFib(), 5; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextFib(), 8; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if val, expected := nextFib(), 13; val != expected {
		t.Fatalf("expected %d, was %d", expected, val)
	}
}
