package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expected string
	}{
		{"no args", []string{}, "\n"},
		{"one arg", []string{"hello"}, "hello\n"},
		{"multiple args", []string{"hello", "world"}, "hello world\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Save current os.Stdout
			originalStdout := os.Stdout

			// Create a new bytes buffer and set it as os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call Echo
			Echo(tc.args)

			// Close the write end of the pipe to finish writing
			w.Close()

			// Read the output
			var buf bytes.Buffer
			buf.ReadFrom(r)

			// Restore os.Stdout
			os.Stdout = originalStdout

			// Check the output
			got := buf.String()
			if got != tc.expected {
				t.Errorf("Echo(%v) = %q, want %q", tc.args, got, tc.expected)
			}
		})
	}
}
