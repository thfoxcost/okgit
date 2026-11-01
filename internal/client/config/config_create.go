package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	internal "okgit/internal"
)

func Create() error {
	fileConfigPath := internal.Path()

	if fileConfigPath == "" {
		return fmt.Errorf("config path is empty")
	}

	if _, err := os.Stat(fileConfigPath); err == nil {
		fmt.Println("Config file already exists")
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("check config file: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(fileConfigPath), 0700); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}

	defaultContent := []byte(`{
	"host": "localhost",
	"port": "3200"
}`)

	if err := os.WriteFile(fileConfigPath, defaultContent, 0600); err != nil {
		return fmt.Errorf("create config file: %w", err)
	}

	fmt.Println("Config file created successfully with default configs")
	return nil
}
