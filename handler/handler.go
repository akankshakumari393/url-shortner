package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var input struct {
	Destination string `json:"destination"`
}

// URLShortener represents the URL shortener application
type URLShortener struct {
}

// NewURLShortener creates a new instance of URLShortener
func NewURLShortener() *URLShortener {
	return &URLShortener{}
}

// WelcomeHandler handles the welcome route
func (u *URLShortener) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the URL Shortener!")
}

// ShortenHandler handles the shorten route
func (u *URLShortener) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handleError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if input.Destination == "" {
		handleError(w, "Missing 'destination' field", http.StatusBadRequest)
		return
	}

	if err = validateUrl(input.Destination); err != nil {
		handleError(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// get the shortURL

	response := map[string]string{"short_url": fmt.Sprintf("http://localhost:8080/r/%s", "")}
	respondJSON(w, response, http.StatusCreated)
}

// RedirectHandler handles the redirect route
func (u *URLShortener) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Add redirection
}
