package utilities

import "testing"

func TestIsPalindrome(t *testing.T) {
	type test struct {
		val  int
		want bool
	}

	tests := []test{
		{1, true},
		{11, true},
		{77, true},
		{303, true},
		{1221, true},
		{1234321, true},
		{10, false},
		{103, false},
		{1223, false},
		{1234521, false},
	}

	for _, tc := range tests {
		if got := IsPalindrome(tc.val); got != tc.want {
			t.Errorf("IsPalindrome(%d) = %v; want %v", tc.val, got, tc.want)
		}
	}
}

func TestReverse(t *testing.T) {
	type test struct {
		val  int
		want int
	}

	tests := []test{
		{1, 1},
		{10, 1},
		{12, 21},
		{33, 33},
		{100, 1},
		{504, 405},
		{1234, 4321},
	}

	for _, tc := range tests {
		if got := Reverse(tc.val); got != tc.want {
			t.Errorf("Reverse(%d) = %d; want %d", tc.val, got, tc.want)
		}
	}
}
