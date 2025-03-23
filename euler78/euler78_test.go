package euler78

import "testing"

func TestEuler78(t *testing.T) {
	answer := 55374

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"DynamicProgramming",
			DynamicProgramming,
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
