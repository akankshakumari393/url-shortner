package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akankshakumari393/url-shortner/urlgenerator"
)

var input struct {
	Destination string `json:"destination"`
}

// URLShortener represents the URL shortener application
type URLShortener struct {
	redisCli urlgenerator.RedisClientProvider
}

// NewURLShortener creates a new instance of URLShortener
func NewURLShortener(redisCli urlgenerator.RedisClientProvider) *URLShortener {
	return &URLShortener{
		redisCli: redisCli,
	}
}

// WelcomeHandler handles the welcome route
func (u *URLShortener) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the URL Shortener!")
}

// ShortenHandler handles the shorten route
func (u *URLShortener) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
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

	shortURL, err := u.redisCli.ShortenURL(ctx, input.Destination)
	if err != nil {
		handleError(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"short_url": fmt.Sprintf("http://localhost:8080/r/%s", shortURL)}
	respondJSON(w, response, http.StatusCreated)
}

// RedirectHandler handles the redirect route
func (u *URLShortener) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	shortURL := mux.Vars(r)["shortURL"]
	url, err := u.redisCli.RedirectToURL(ctx, shortURL)
	if err != nil {
		handleError(w, "URL Not Found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}
