package handlers

import (
	"net/http"

	"github.com/prasadsurase/developer-portfolio-in-go/config"
	"github.com/prasadsurase/developer-portfolio-in-go/models"
	"github.com/prasadsurase/developer-portfolio-in-go/pkg/render"
)

// Repository is the repository type
type Repository struct {
	AppConfig *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["name"] = "Prasad Surase"
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
