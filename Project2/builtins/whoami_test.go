package builtins

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"testing"
)

type MockUserRetriever struct {
	User *user.User
	Err  error
}

func (m *MockUserRetriever) Current() (*user.User, error) {
	return m.User, m.Err
}

func TestWhoamiSuccess(t *testing.T) {
	mockUser := &user.User{Username: "testuser"}
	retriever := &MockUserRetriever{User: mockUser, Err: nil}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Whoami(retriever)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if output != mockUser.Username+"\n" {
		t.Errorf("Expected username %s, got %s", mockUser.Username, output)
	}
}

func TestWhoamiError(t *testing.T) {
	retriever := &MockUserRetriever{Err: fmt.Errorf("error retrieving user")}

	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	Whoami(retriever)

	w.Close()
	os.Stderr = oldStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("whoami error")) {
		t.Errorf("Expected error output, got %s", output)
	}
}
