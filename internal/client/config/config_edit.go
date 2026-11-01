package config

import (
	"fmt"
	"os"
	"os/exec"

	internal "okgit/internal"
)

func Edit() error {
	fileConfigPath := internal.Path()

	if fileConfigPath == "" {
		return fmt.Errorf("config path is empty")
	}

	cmd := exec.Command("nano", fileConfigPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
