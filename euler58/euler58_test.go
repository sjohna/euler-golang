package euler58

import "testing"

func TestEuler58(t *testing.T) {
	answer := 26241

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
