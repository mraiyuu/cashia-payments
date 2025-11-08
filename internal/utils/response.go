package utils

import (
	"net/http"

	"github.com/bytedance/sonic"
)

type jsonresponsesuccess struct {
	Success string      `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonresponseerror struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data,omitempty"`
}

func WriteJSONSuccess(w http.ResponseWriter, message string, statusCode int, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := jsonresponsesuccess{
		Success: message,
	}

	// Only add data if provided
	if len(data) > 0 && data[0] != nil {
		response.Data = data[0]
	}

	// Use sonic for fast encoding
	if err := sonic.ConfigDefault.NewEncoder(w).Encode(response); err != nil {
		// Fallback to plain error if encoding fails
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func WriteJSONError(w http.ResponseWriter, message string, statusCode int, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := jsonresponseerror{
		Error: message,
	}

	// Only add data if provided
	if len(data) > 0 && data[0] != nil {
		response.Data = data[0]
	}

	// Use sonic for fast encoding
	if err := sonic.ConfigDefault.NewEncoder(w).Encode(response); err != nil {
		// Fallback to plain error if encoding fails
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}