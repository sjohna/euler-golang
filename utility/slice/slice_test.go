package slice

import (
	"euler/utility/types"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type testCase[T types.Number] struct {
		slice []T
		want  T
	}
	intTests := []testCase[int]{
		{
			[]int{},
			0,
		},
		{
			[]int{7},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			15,
		},
	}
	for i, tt := range intTests {
		t.Run(fmt.Sprintf("IntSum%d", i+1), func(t *testing.T) {
			if got := Sum(tt.slice); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}

	floatTests := []testCase[float64]{
		{
			[]float64{},
			0.0,
		},
		{
			[]float64{1.0, 2.5},
			3.5,
		},
	}
	for i, tt := range floatTests {
		t.Run(fmt.Sprintf("FloatSum%d", i+1), func(t *testing.T) {
			if got := Sum(tt.slice); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	type testCase[T types.Number] struct {
		slice []T
		want  T
	}
	intTests := []testCase[int]{
		{
			[]int{},
			1,
		},
		{
			[]int{7},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			120,
		},
		{
			[]int{1, 2, 3, 4, -5},
			-120,
		},
		{
			[]int{1, 2, 3, 4, 5, 0},
			0,
		},
	}
	for i, tt := range intTests {
		t.Run(fmt.Sprintf("IntProduct%d", i+1), func(t *testing.T) {
			if got := Product(tt.slice); got != tt.want {
				t.Errorf("Product() = %v, want %v", got, tt.want)
			}
		})
	}

	floatTests := []testCase[float64]{
		{
			[]float64{},
			1.0,
		},
		{
			[]float64{1.0, 2.5},
			2.5,
		},
		{
			[]float64{1.0, 2.5, 3.754456, 0.0},
			0.0,
		},
	}
	for i, tt := range floatTests {
		t.Run(fmt.Sprintf("FloatProduct%d", i+1), func(t *testing.T) {
			if got := Product(tt.slice); got != tt.want {
				t.Errorf("Product() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	type testCase[T comparable] struct {
		slice []T
		want  bool
	}

	intTests := []testCase[int]{
		{
			[]int{},
			true,
		},
		{
			[]int{7},
			true,
		},
		{
			[]int{1, 2},
			false,
		},
		{
			[]int{2, 2},
			true,
		},
		{
			[]int{1, 2, 3},
			false,
		},
		{
			[]int{7, 6, 7},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 9, 0, 9, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 9, 0, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			false,
		},
	}

	for i, tt := range intTests {
		t.Run(fmt.Sprintf("IsPalindrome%d", i+1), func(t *testing.T) {
			if got := IsPalindrome(tt.slice); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
