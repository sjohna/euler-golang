package utilities

import (
	"reflect"
	"testing"
)

func TestFibonacciGeneratorFirstValues(t *testing.T) {
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	got := TakeN(FibonacciSequence(), 10)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
