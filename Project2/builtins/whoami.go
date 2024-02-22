package builtins

import (
	"os/user"
)

// Whoami returns the username of the current user.
func Whoami() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err // Return the error to the caller
	}
	return user.Username, nil
}
