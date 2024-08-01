package config

import "html/template"

// AppConfig holds the configuration for the application.
type AppConfig struct {
	UseCache       bool
	TemplatesCache map[string]*template.Template
}
