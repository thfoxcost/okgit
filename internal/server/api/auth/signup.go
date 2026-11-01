package auth

import (
	"encoding/json"
	"net/http"
)

type SignupData struct {
	Username string `json:"username"`
	Email    string `json:email`
	Password string `json:"password"`
}

func SingupHandlerR(w http.ResponseWriter, r *http.Request) {
	var data SignupData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "\033[31mInvalid request body\033[0m", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Account created successfully"))
}
