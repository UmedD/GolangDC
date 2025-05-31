package http_rest

import (
	"encoding/json"
	"net/http"
)

type Err struct {
	Error string `json:"error"`
}

func handleError(w http.ResponseWriter, statusCode int, errStr string) {
	w.WriteHeader(statusCode)
	//w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		Err{errStr},
	)
}
