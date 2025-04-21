package euler63

import "testing"

func TestEuler63(t *testing.T) {
	answer := 49

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
