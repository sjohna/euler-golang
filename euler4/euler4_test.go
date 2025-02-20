package euler4

import "testing"

func TestEuler4(t *testing.T) {
	answer := 906609

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Loops",
			Loops,
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
