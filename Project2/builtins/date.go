package builtins

import (
	"fmt"
	"time"
)

// Date prints the current date and time.
func Date() {
	fmt.Println(time.Now().Format(time.RFC1123))
}
