package builtins

import (
	"os"
)

// Pwd returns the current working directory.
func Pwd() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err // Return the error to the caller
	}
	return dir, nil
}
