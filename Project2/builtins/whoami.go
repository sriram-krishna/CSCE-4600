package builtins

import (
	"fmt"
	"os/user"
)

// Whoami prints the current user's username.
func Whoami() error {
	currentUser, err := user.Current()
	if err != nil {
		return err // Return the error to be handled by the caller
	}
	fmt.Println(currentUser.Username)
	return nil
}
