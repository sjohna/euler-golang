package utility

import (
	"fmt"
	"testing"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			1,
			1,
			11,
		},
		{
			2,
			1,
			21,
		},
		{
			12,
			1,
			121,
		},
		{
			1,
			21,
			121,
		},
		{
			1234,
			56789,
			123456789,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Concat(%d,%d)", tt.a, tt.b), func(t *testing.T) {
			if got := Concat(tt.a, tt.b); got != tt.want {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
