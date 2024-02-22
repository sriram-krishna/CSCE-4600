package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	// Define test cases
	tests := []struct {
		name    string
		input   string
		wantOut string
		wantErr string
	}{
		{
			name:    "exit command",
			input:   "exit\n",
			wantOut: "exiting gracefully...\n",
		},
		{
			name:    "echo command",
			input:   "echo Hello, World!\nexit\n",
			wantOut: "Hello, World!\n",
		},
		// Additional test cases for pwd, whoami, etc.
		// Assuming mock or direct calls to builtins for simplicity
		{
			name:    "pwd command",
			input:   "pwd\nexit\n",
			wantOut: "/your/expected/path\n", // Adjust this to match your expected or mocked output
		},
		{
			name:    "whoami command",
			input:   "whoami\nexit\n",
			wantOut: "testuser\n", // Adjust this to your mock or expected output
		},
		// Example error scenario for cd command to a nonexistent directory
		{
			name:    "cd to invalid directory",
			input:   "cd /nonexistent\nexit\n",
			wantErr: "no such file or directory", // This is illustrative; adjust based on actual error handling
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}
			exit := make(chan struct{}, 2)
			input := strings.NewReader(tt.input)

			go runLoop(input, w, errW, exit)
			time.Sleep(100 * time.Millisecond) // Adjust timing as necessary
			exit <- struct{}{}

			if tt.wantOut != "" && !strings.Contains(w.String(), tt.wantOut) {
				t.Errorf("Output = %v, want %v", w.String(), tt.wantOut)
			}
			if tt.wantErr != "" && !strings.Contains(errW.String(), tt.wantErr) {
				t.Errorf("Error Output = %v, want %v", errW.String(), tt.wantErr)
			}
		})
	}
}
