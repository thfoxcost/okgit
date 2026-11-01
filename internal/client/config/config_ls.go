package config

import (
	"fmt"
	"os"

	internal "okgit/internal"
)

func Ls() error {
	fileConfigPath := internal.Path()
	if fileConfigPath == "" {
		fmt.Println("config file path does not exist")
		return nil
	}

	content, err := os.ReadFile(fileConfigPath)
	if err != nil {
		return nil
	}
	fmt.Println(string(content))
	return nil
}
