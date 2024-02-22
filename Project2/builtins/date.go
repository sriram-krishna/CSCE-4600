package builtins

import (
	"time"
)

// Date returns the current date and time in a predefined format.
func Date() string {
	return time.Now().Format("2006-01-02 15:04:05 MST")
}
