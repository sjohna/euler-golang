package euler32

import "testing"

func TestEuler32(t *testing.T) {
	answer := 45228

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
