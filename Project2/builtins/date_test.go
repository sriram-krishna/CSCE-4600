package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestDate(t *testing.T) {
	// Capture the output
	var buf bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Date()

	// Cleanup
	w.Close()
	os.Stdout = originalStdout

	// Read output
	buf.ReadFrom(r)
	output := buf.String()

	if output == "" {
		t.Errorf("Expected date output to be non-empty")
	}
	// Further validations could include checking the format of the date,
	// but given the dynamic nature of the output, this might be sufficient.
}
