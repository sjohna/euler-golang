package utility

import (
	"testing"
)

func TestRational_Plus(t *testing.T) {
	type test struct {
		a        Rational
		b        Rational
		expected Rational
	}

	tests := []test{
		{
			Rational{1, 1},
			Rational{1, 1},
			Rational{2, 1},
		},
		{
			Rational{0, 1},
			Rational{1, 1},
			Rational{1, 1},
		},
		{
			Rational{1, 2},
			Rational{1, 2},
			Rational{1, 1},
		},
		{
			Rational{1, 1},
			Rational{1, 2},
			Rational{3, 2},
		},
		{
			Rational{1, 2},
			Rational{2, 5},
			Rational{9, 10},
		},
		{
			Rational{-1, 2},
			Rational{1, 2},
			Rational{0, 1},
		},
	}

	for _, test := range tests {
		result := test.a.Plus(test.b)
		if !result.Equals(test.expected) {
			t.Errorf("%v + %v = %v, but expected %v", test.a.ToString(), test.b.ToString(), result.ToString(), test.expected.ToString())
		}
	}
}
