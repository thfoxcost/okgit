package internal

import (
	"os"
)

func Path() string {
	return os.Getenv("OKGIT_CONFIG_PATH")
}
