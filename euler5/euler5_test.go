package euler5

import "testing"

func TestEuler5(t *testing.T) {
	answer := 232792560

	type test struct {
		name     string
		function func() int
	}

	tests := []test{
		{
			"LoopLCMFunction",
			LoopLCMFunction,
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
