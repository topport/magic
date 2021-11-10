package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"

	"github.com/topport/magic/cmd/magic/command"
)

func main() {
	if err := command.Execute(); err != nil {
		fatalf("%s %v\n\n", color.RedString("Error:"), err)
	}
}

func fatalf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	// revive:disable-next-line:deep-exit different impls exit at different points
	os.Exit(1)
}
