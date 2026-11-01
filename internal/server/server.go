package server

import (
	"net/http"

	"okgit/internal/server/api"
	"okgit/internal/server/api/auth"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/status", api.Status)

	// authentication
	mux.HandleFunc("POST /api/auth/signup", auth.SingupHandlerR)

	return mux

}
