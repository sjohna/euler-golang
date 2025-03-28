package euler92

import "testing"

func TestEuler92(t *testing.T) {
	answer := 8581146

	type test struct {
		name     string
		function func() int
	}

	// not testing the dynamic programming implementation, because it takes ~45 seconds
	tests := []test{
		{
			"Recursively",
			Recursively,
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
