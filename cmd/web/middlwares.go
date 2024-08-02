package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WritetoConsole prints the request url
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hit url: %v\n", r.URL)
		next.ServeHTTP(rw, r)
	})
}

// NoSurf handles the CSRF protection to all non-GET requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoader loads and handles the session for every request
func SessionLoader(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
