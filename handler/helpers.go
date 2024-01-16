package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// handleError sends an error response with the specified message and status code
func handleError(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]string{"error": message}
	respondJSON(w, response, statusCode)
}

// respondJSON sends a JSON response with the specified data and status code
func respondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func validateUrl(inputUrl string) error {
	_, err := url.ParseRequestURI(inputUrl)
	if err != nil {
		return err
	}
	return nil
}
