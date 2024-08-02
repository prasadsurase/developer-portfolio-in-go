package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prasadsurase/developer-portfolio-in-go/config"
	"github.com/prasadsurase/developer-portfolio-in-go/pkg/handlers"
	"github.com/prasadsurase/developer-portfolio-in-go/pkg/render"
)

const serverPort = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.UseCache = false
	app.TemplatesCache = tc
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", serverPort)

	srv := &http.Server{
		Addr:    serverPort,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
