package euler2

import "testing"

func TestEuler2(t *testing.T) {
	answer := 4613732

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"GeneratorWithLoop",
			GeneratorWithLoop,
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
