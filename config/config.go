package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the configuration for the application.
type AppConfig struct {
	UseCache       bool
	TemplatesCache map[string]*template.Template
	InProduction   bool
	SessionManager *scs.SessionManager
}
