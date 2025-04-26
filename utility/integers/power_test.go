package integers

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	type testCase struct {
		base int
		exp  int
		want int
	}
	tests := []testCase{
		{
			1, 1, 1,
		},
		{
			2, 1, 2,
		},
		{
			2, 3, 8,
		},
		{
			5, 5, 3125,
		},
		{
			2, 16, 65536,
		},
		{
			2, 31, 2147483648,
		},
		{
			2, 32, 4294967296,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Pow(%d,%d)", tt.base, tt.exp), func(t *testing.T) {
			if got := Pow(tt.base, tt.exp); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPowMod(t *testing.T) {
	type testCase struct {
		base int
		exp  int
		mod  int
		want int
	}
	tests := []testCase{
		{
			1, 1, 10, 1,
		},
		{
			2, 1, 10, 2,
		},
		{
			2, 3, 10, 8,
		},
		{
			5, 5, 10000, 3125,
		},
		{
			2, 16, 1000000, 65536,
		},
		{
			2, 31, 10000000000, 2147483648,
		},
		{
			2, 32, 10000000000, 4294967296,
		},
		{
			5, 5, 100, 25,
		},
		{
			2, 31, 19, 2147483648 % 19,
		},
		{
			5, 123456789, 4, 1,
		},
		{
			5, 123456789, 6, 5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("PowMod(%d,%d,%d)", tt.base, tt.exp, tt.mod), func(t *testing.T) {
			if got := PowMod(tt.base, tt.exp, tt.mod); got != tt.want {
				t.Errorf("PowMod() = %v, want %v", got, tt.want)
			}
		})
	}
}
