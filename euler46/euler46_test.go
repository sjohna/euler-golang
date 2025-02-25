package euler46

import "testing"

func TestEuler46(t *testing.T) {
	answer := 5777

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Heap",
			Heap,
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
