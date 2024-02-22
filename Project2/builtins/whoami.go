package builtins

import (
	"fmt"
	"os"
	"os/user"
)

// UserRetriever defines an interface for retrieving user details.
type UserRetriever interface {
	Current() (*user.User, error)
}

// RealUserRetriever retrieves real user information using the os/user package.
type RealUserRetriever struct{}

func (r *RealUserRetriever) Current() (*user.User, error) {
	return user.Current()
}

// Whoami prints the current user's username using the provided UserRetriever.
func Whoami(retriever UserRetriever) {
	currentUser, err := retriever.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "whoami error: %v\n", err)
		return
	}
	fmt.Println(currentUser.Username)
}
