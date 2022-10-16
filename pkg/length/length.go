package length

import "fmt"

// Print length of first gavin args
func Print(args []string) {
	if len(args) > 0 {
		fmt.Println(len(args[0]))
	}
}
