package main

import (
	"log"
	"net/http"

	internal "okgit/internal"
	"okgit/internal/server"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	_, minEndpoint, err := internal.Endpoint()
	if err != nil {
		return
	}

	log.Printf("okgit server listening on %s", minEndpoint)

	if err := http.ListenAndServe(minEndpoint, server.New()); err != nil {
		log.Fatal(err)
	}
}
