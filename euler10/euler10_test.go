package euler10

import "testing"

func TestEuler10(t *testing.T) {
	answer := 142913828922

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
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
