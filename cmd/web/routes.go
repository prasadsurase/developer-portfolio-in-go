package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/prasadsurase/developer-portfolio-in-go/config"
	"github.com/prasadsurase/developer-portfolio-in-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
