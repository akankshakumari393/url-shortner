package main

import (
	"log"
	"net/http"
	"time"

	"github.com/akankshakumari393/url-shortner/handler"
	"github.com/akankshakumari393/url-shortner/middleware"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	urlShortener := handler.NewURLShortener()

	// Define routes
	router.HandleFunc("/", urlShortener.WelcomeHandler).Methods("GET")
	router.HandleFunc("/shortcode", urlShortener.ShortenHandler).Methods("PUT")
	router.HandleFunc("/r/{shortURL}", urlShortener.RedirectHandler).Methods("GET")

	router.Use(middleware.LoggingMiddleware)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
