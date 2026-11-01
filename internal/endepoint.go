package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func Endpoint() (string, string, error) {
	data, err := os.ReadFile(Path())
	if err != nil {
		fmt.Print(err)
		return "", "", nil
	}

	var config ServerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return "", "", err
	}

	endpoint := "http://" + config.Host + ":" + config.Port
	minEndpoint := config.Host + ":" + config.Port

	return endpoint, minEndpoint, nil
}
