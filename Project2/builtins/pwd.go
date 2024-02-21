package builtins

import (
	"fmt"
	"os"
)

// Pwd prints the current working directory.
func Pwd() {
	if dir, err := os.Getwd(); err == nil {
		fmt.Println(dir)
	} else {
		fmt.Fprintf(os.Stderr, "pwd error: %v\n", err)
	}
}
