package utility

import (
	"testing"
)

func TestGenerator_SkipWhile(t *testing.T) {
	// skip some
	g := NaturalNumbers()
	if value := g.SkipWhile(LessThan(1000)).NextValue(); value != 1000 {
		t.Errorf("expected 1000, got %d", value)
	}

	// don't skip any
	g = NaturalNumbers()
	if value := g.SkipWhile(LessThan(-10)).NextValue(); value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}
