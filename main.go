package main

import (
	"fmt"
	"os"

	"github.com/bugfire/qnapcc/cmd"
)

func main() {
	var exitCode int
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
	}
	os.Exit(exitCode)
}
