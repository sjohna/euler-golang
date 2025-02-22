package utility

import (
	"reflect"
	"testing"
)

func TestKthPermutationOfNItems(t *testing.T) {
	type test struct {
		N        int
		K        int
		expected []int
	}

	tests := []test{
		{3, 1, []int{0, 1, 2}},
		{3, 2, []int{0, 2, 1}},
		{3, 3, []int{1, 0, 2}},
		{3, 4, []int{1, 2, 0}},
		{3, 5, []int{2, 0, 1}},
		{3, 6, []int{2, 1, 0}},
		{4, 1, []int{0, 1, 2, 3}},
		{4, 8, []int{1, 0, 3, 2}},
	}

	for _, test := range tests {
		actual := KthPermutationOfNItems(test.N, test.K)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%dth permutation of %d items: expected %v, got %v", test.K, test.N, test.expected, actual)
		}
	}
}
