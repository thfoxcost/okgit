package config

import (
	"errors"
	"fmt"
	"os"

	internal "okgit/internal"
)

func Rm() error {
	fileConfigPath := internal.Path()

	_, err := os.Stat(fileConfigPath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("config file does not exist")
	}
	if err != nil {
		return fmt.Errorf("check config file: %w", err)
	}

	if err := os.Remove(fileConfigPath); err != nil {
		return fmt.Errorf("delete config file: %w", err)
	}

	fmt.Println("Config file deleted successfully")
	return nil
}
