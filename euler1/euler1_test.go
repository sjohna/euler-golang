package euler1

import "testing"

func TestEuler1(t *testing.T) {
	answer := 233168

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Loop",
			Loop,
		},
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
