package euler33

import "testing"

func TestEuler33(t *testing.T) {
	answer := 100

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Loop",
			Loop,
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
