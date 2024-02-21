package builtins

import (
	"bytes"
	"testing"
)

func TestClear(t *testing.T) {
	var buf bytes.Buffer
	Clear(&buf)
	expected := "\033[H\033[2J"
	if got := buf.String(); got != expected {
		t.Errorf("Clear() = %q, want %q", got, expected)
	}
}
