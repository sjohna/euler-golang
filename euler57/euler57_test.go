package euler57

import "testing"

func TestEuler57(t *testing.T) {
	answer := 153

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Iterate",
			Iterate,
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
