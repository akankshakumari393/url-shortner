package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/akankshakumari393/url-shortner/urlgenerator"
)

func TestURLShortener(t *testing.T) {
	// Create a Redis client for testing
	redisClient := urlgenerator.NewMockRedisClient()

	// Create a URLShortener instance
	urlShortener := NewURLShortener(redisClient)

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Attach the handlers to the router
	router.HandleFunc("/", urlShortener.WelcomeHandler).Methods("GET")
	router.HandleFunc("/shorten", urlShortener.ShortenHandler).Methods("POST")
	router.HandleFunc("/r/{shortURL}", urlShortener.RedirectHandler).Methods("GET")

	// Test cases
	t.Run("WelcomeHandler", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(urlShortener.WelcomeHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Welcome to the URL Shortener!")
	})

	t.Run("ShortenHandler", func(t *testing.T) {
		inputJSON := `{"destination": "https://example.com"}`
		req, err := http.NewRequest("PUT", "/shorten", bytes.NewBufferString(inputJSON))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(urlShortener.ShortenHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		var response map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, response["short_url"], "http://localhost:8080/r/mockShortURL")
	})

	t.Run("RedirectHandler", func(t *testing.T) {
		// Assuming a short URL "abcdef12" is generated and stored in Redis
		// before running this test
		shortURL := "mockShortURL"
		req, err := http.NewRequest("GET", "/r/"+shortURL, nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(urlShortener.RedirectHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusSeeOther, rr.Code)
		assert.Contains(t, rr.Header().Get("Location"), "https://example.com")
	})
}
