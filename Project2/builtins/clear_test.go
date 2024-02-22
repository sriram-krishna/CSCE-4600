package builtins_test

import (
	"testing"

	"github.com/sriram-krishna/CSCE-4600/Project2/builtins"
)

func TestClear(t *testing.T) {
	want := "\033[H\033[2J"
	if got := builtins.Clear(); got != want {
		t.Errorf("Clear() = %v, want %v", got, want)
	}
}
