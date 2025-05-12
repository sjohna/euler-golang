package euler60

import (
	"fmt"
	"testing"
)

func TestFindCliqueSum(t *testing.T) {
	tests := []struct {
		size int
		want int
	}{
		{
			4, 792,
		},
		{
			5, 26033,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("FindCliqueSum(%d)", tt.size), func(t *testing.T) {
			if got := FindCliqueSum(tt.size); got != tt.want {
				t.Errorf("FindCliqueSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
