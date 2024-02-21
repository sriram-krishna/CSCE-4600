package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestPwd(t *testing.T) {
	// Expected current directory
	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Capture the output
	var buf bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Pwd()

	// Cleanup
	w.Close()
	os.Stdout = originalStdout

	// Read output
	buf.ReadFrom(r)
	output := buf.String()

	// The output will have a newline at the end, so we trim it before comparison
	if output != expected+"\n" {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
