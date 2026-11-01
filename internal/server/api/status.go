package api

import (
	"encoding/json"
	"net/http"
)

type Respone struct {
	Message string `json:message`
	Status  string `json:status`
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Respone{
		Message: "hello from the paradise",
		Status:  "ok",
	})

}
