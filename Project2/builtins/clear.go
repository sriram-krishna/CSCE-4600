package builtins

import (
	"fmt"
	"io"
)

// Clear sends ANSI escape codes to clear the terminal screen.
func Clear(w io.Writer) {
	fmt.Fprint(w, "\033[H\033[2J")
}
