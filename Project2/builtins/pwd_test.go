package builtins_test

import (
	"testing"

	"github.com/sriram-krishna/CSCE-4600/Project2/builtins"
)

func TestPwd(t *testing.T) {
	got, err := builtins.Pwd()
	if err != nil {
		t.Errorf("Pwd() error = %v, wantErr false", err)
	}
	if got == "" {
		t.Errorf("Pwd() returned an empty string, expected current working directory")
	}
}
