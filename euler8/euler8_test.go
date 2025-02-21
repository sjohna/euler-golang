package euler8

import "testing"

func TestEuler8(t *testing.T) {
	answer := 23514624000

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"Naive",
			Naive,
		},
		{
			"Generator",
			Generator,
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
