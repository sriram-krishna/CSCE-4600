package builtins_test

import (
	"testing"

	"github.com/sriram-krishna/CSCE-4600/Project2/builtins"
)

func TestWhoami(t *testing.T) {
	got, err := builtins.Whoami()
	if err != nil {
		t.Errorf("Whoami() error = %v, wantErr false", err)
	}
	if got == "" {
		t.Errorf("Whoami() returned an empty string, expected a username")
	}
}
