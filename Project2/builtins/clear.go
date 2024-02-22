package builtins

// Clear returns the ANSI escape sequence for clearing the terminal screen.
func Clear() string {
	return "\033[H\033[2J"
}
