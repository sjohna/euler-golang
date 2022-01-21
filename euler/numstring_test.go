package euler

import "testing"

func TestIsPalindrome(t *testing.T) {
	if val, expected := 1, true; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 11, true; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 77, true; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 303, true; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 1221, true; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 1234321, true; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 10, false; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 103, false; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 1223, false; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}

	if val, expected := 1234351, false; IsPalindrome(val) != expected {
		t.Fatalf("for %d: expected %v, was %v", val, expected, !expected)
	}
}
