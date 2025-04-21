package euler62

import (
	"fmt"
	"testing"
)

func TestSortDigits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{0, 0},
		{12, 21},
		{123, 321},
		{58733129, 98753321},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("SortDigits(%d)", tt.n), func(t *testing.T) {
			if got := SortDigits(tt.n); got != tt.want {
				t.Errorf("SortDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEuler62(t *testing.T) {
	answer := 127035954683

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"SolveIt",
			SolveIt,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := tc.function(); result != answer {
				t.Errorf("expected %v, got %v", answer, result)
			}
		})
	}
}
