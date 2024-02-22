package builtins

import (
	"strings"
)

// Echo concatenates the input arguments with a space and returns the result.
func Echo(args ...string) string {
	return strings.Join(args, " ")
}
