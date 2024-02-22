package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}

func TestRunLoopExit(t *testing.T) {
	t.Parallel()
	commandInput := strings.NewReader("exit\n")
	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	exit := make(chan struct{}, 2)

	go runLoop(commandInput, w, errW, exit)
	time.Sleep(10 * time.Millisecond) // Adjust timing as necessary
	exit <- struct{}{}

	require.Contains(t, w.String(), "exiting gracefully")
}

func TestEchoCommand(t *testing.T) {
	t.Parallel()
	commandInput := strings.NewReader("echo Hello, World!\nexit\n")
	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	exit := make(chan struct{}, 2)

	go runLoop(commandInput, w, errW, exit)
	time.Sleep(10 * time.Millisecond)
	exit <- struct{}{}

	require.Contains(t, w.String(), "Hello, World!")
	require.Empty(t, errW.String())
}

func TestWhoamiCommand(t *testing.T) {
	// Assuming "whoami" just prints a fixed string for simplicity
	t.Parallel()
	commandInput := strings.NewReader("whoami\nexit\n")
	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	exit := make(chan struct{}, 2)

	go runLoop(commandInput, w, errW, exit)
	time.Sleep(10 * time.Millisecond)
	exit <- struct{}{}

	// Check for the expected username in the output
	// In a real scenario, you might mock the output or check for a real username
	require.Contains(t, w.String(), "testuser") // Adjust according to your mock setup
	require.Empty(t, errW.String())
}

func TestInvalidCommand(t *testing.T) {
	t.Parallel()
	commandInput := strings.NewReader("invalidcmd\nexit\n")
	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	exit := make(chan struct{}, 2)

	go runLoop(commandInput, w, errW, exit)
	time.Sleep(10 * time.Millisecond)
	exit <- struct{}{}

	require.NotEmpty(t, errW.String())
}

func TestPwdCommand(t *testing.T) {
	t.Parallel()
	// Simulate entering the "pwd" command and then "exit"
	commandInput := strings.NewReader("pwd\nexit\n")
	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	exit := make(chan struct{}, 2)

	go runLoop(commandInput, w, errW, exit)
	time.Sleep(10 * time.Millisecond) // Adjust timing as necessary
	exit <- struct{}{}

	// The actual output will depend on the test environment's current directory
	// For simplicity, we're just checking if the output is not empty
	require.NotEmpty(t, w.String(), "Expected non-empty output from pwd command")
}

func TestClearCommand(t *testing.T) {
	t.Parallel()
	// "clear" command doesn't produce a direct output that's easily tested,
	// so we'll focus on it not producing an error and being recognized.
	commandInput := strings.NewReader("clear\nexit\n")
	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	exit := make(chan struct{}, 2)

	go runLoop(commandInput, w, errW, exit)
	time.Sleep(10 * time.Millisecond)
	exit <- struct{}{}

	// Since "clear" typically clears the terminal, which isn't visible in test output,
	// the absence of errors can be a sufficient check.
	require.Empty(t, errW.String(), "Expected no error output from clear command")
}

func TestErrorInExecuteCommand(t *testing.T) {
	t.Parallel()
	// Simulate entering an invalid command to trigger an error in executeCommand
	commandInput := strings.NewReader("invalid_command\nexit\n")
	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	exit := make(chan struct{}, 2)

	go runLoop(commandInput, w, errW, exit)
	time.Sleep(10 * time.Millisecond)
	exit <- struct{}{}

	// Check for an error message indicating the command was not found or executed
	require.NotEmpty(t, errW.String(), "Expected an error message for an invalid command")
}
