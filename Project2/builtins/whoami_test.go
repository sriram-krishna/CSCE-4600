package builtins

import (
	"bytes"
	"os"
	"os/user"
	"testing"
)

func TestWhoami(t *testing.T) {
	// Redirect stdout to capture the output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Whoami()

	// Close the write end and restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Read and verify the output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	currentUser, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}

	expected := currentUser.Username + "\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
