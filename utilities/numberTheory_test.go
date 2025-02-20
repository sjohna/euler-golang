package utilities

import "testing"

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
