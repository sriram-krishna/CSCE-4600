package builtins

import (
	"fmt"
	"strings"
)

// Echo prints the arguments passed to it
func Echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}
