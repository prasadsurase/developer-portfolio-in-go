package handlers

import (
	"net/http"

	"github.com/prasadsurase/developer-portfolio-in-go/pkg/render"
)

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.tmpl")
}
