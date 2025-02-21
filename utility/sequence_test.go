package utility

import (
	"reflect"
	"testing"
)

func TestFibonacciGeneratorFirstValues(t *testing.T) {
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	got := FibonacciSequence().Take(10).ToSlice()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
