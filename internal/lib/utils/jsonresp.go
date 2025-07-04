// File: internal/lib/utils/jsonresp.go
// Purpose: Shared helpers for writing JSON responses and handling HTTP errors.

package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteJSON sends a successful JSON response.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := JSONResponse{
		Status: "ok",
		Data:   data,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

// WriteError sends an error JSON response with custom message and status.
func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := JSONResponse{
		Status:  "error",
		Message: message,
	}

	_ = json.NewEncoder(w).Encode(resp)
}
