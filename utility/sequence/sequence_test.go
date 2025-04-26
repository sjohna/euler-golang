package sequence

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	got := Fibonacci().Take(10).ToSlice()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestGeneralizedFibonacci(t *testing.T) {
	// regular fibonacci
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	got := GeneralizedFibonacci(1, 1).Take(10).ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	// tribonacci
	expected = []int{1, 1, 1, 3, 5, 9, 17, 31, 57, 105}
	got = GeneralizedFibonacci(1, 1, 1).Take(10).ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	// negative fibonacci
	expected = []int{-1, -1, -2, -3, -5, -8, -13, -21, -34, -55}
	got = GeneralizedFibonacci(-1, -1).Take(10).ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
