package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

var (
	version = "dev"
	commit  = ""
)

func main() {
	if slices.Contains(os.Args, "-v") || slices.Contains(os.Args, "--version") {
		fmt.Printf("Version: %s\nCommit: %s\n", version, commit)
		os.Exit(0)
	}

	fmt.Println("Hello, World!")
}
