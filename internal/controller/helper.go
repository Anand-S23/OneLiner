package controller

import (
	"encoding/json"
	"net/http"
)

func InternalServerError(w http.ResponseWriter) error {
    errMsg := map[string]string {
        "error": "Internal server error, please try again",
    }

    return WriteJSON(w, http.StatusInternalServerError, errMsg)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
