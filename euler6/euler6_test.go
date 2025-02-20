package euler6

import "testing"

func TestEuler5(t *testing.T) {
	answer := 25164150

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Generators",
			Generators,
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
