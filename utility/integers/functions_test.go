package integers

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	type test struct {
		a    int
		b    int
		want int
	}

	tests := []test{
		{1, 1, 1},
		{1, 7, 1},
		{2, 3, 1},
		{2, 2, 2},
		{7, 7, 7},
		{2, 4, 2},
		{4, 2, 2},
		{6, 10, 2},
		{35, 1, 1},
		{2 * 3 * 5 * 5 * 7 * 13, 2 * 5 * 7 * 7 * 11 * 13, 2 * 5 * 7 * 13},
	}

	for _, tc := range tests {
		if got := GCD(tc.a, tc.b); got != tc.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestLCM(t *testing.T) {
	type test struct {
		a    int
		b    int
		want int
	}

	tests := []test{
		{1, 1, 1},
		{2, 2, 2},
		{2, 3, 6},
		{1, 10, 10},
		{4, 6, 12},
		{6, 4, 12},
		{5, 7, 35},
		{2 * 2 * 2 * 3 * 5 * 7, 2 * 2 * 3 * 3 * 7, 2 * 2 * 2 * 3 * 3 * 5 * 7},
	}

	for _, tc := range tests {
		if got := LCM(tc.a, tc.b); got != tc.want {
			t.Errorf("LCM(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			1,
			1,
			11,
		},
		{
			2,
			1,
			21,
		},
		{
			12,
			1,
			121,
		},
		{
			1,
			21,
			121,
		},
		{
			1234,
			56789,
			123456789,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Concat(%d,%d)", tt.a, tt.b), func(t *testing.T) {
			if got := Concat(tt.a, tt.b); got != tt.want {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{11211, 11211},
		{9876543210, 123456789},
	}

	for _, tc := range tests {
		if got := Reverse(tc.val); got != tc.want {
			t.Errorf("Reverse(%d) = %d; want %d", tc.val, got, tc.want)
		}
	}
}
