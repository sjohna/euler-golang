package euler40

import "testing"

func TestEuler40(t *testing.T) {
	answer := 210

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Basic",
			Basic,
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
