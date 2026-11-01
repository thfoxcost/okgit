package main

import (
	"fmt"
	"os"

	"okgit/cmd"
)

func main() {
	if os.Getenv("OKGIT_CONFIG_PATH") == "" {
		fmt.Fprintln(os.Stderr, "Error: OKGIT_CONFIG_PATH is not set")
		os.Exit(1)
	}

	cmd.Execute()
}
