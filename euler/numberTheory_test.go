package euler

import "testing"

func TestGCD(t *testing.T) {
	if expected, val := 1, GCD(1, 1); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 1, GCD(1, 7); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 1, GCD(2, 3); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 2, GCD(2, 2); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 7, GCD(7, 7); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 2, GCD(2, 4); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 2, GCD(6, 10); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 1, GCD(35, 1); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}
}

func TestLCM(t *testing.T) {
	if expected, val := 1, LCM(1, 1); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 2, LCM(2, 2); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 6, LCM(2, 3); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 10, LCM(1, 10); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 12, LCM(4, 6); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}

	if expected, val := 35, LCM(5, 7); expected != val {
		t.Fatalf("expected %d, was %d", expected, val)
	}
}
