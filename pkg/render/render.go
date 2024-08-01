package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/prasadsurase/developer-portfolio-in-go/config"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate is used to render the template
func RenderTemplate(rw http.ResponseWriter, tmplName string) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplatesCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("could not create template cache")
		}
	}

	t, ok := tc[tmplName]
	if !ok {
		log.Fatal("could not get template from cache")
	}

	buf := new(bytes.Buffer)
	t.Execute(buf, nil)

	_, err = buf.WriteTo(rw)
	if err != nil {
		fmt.Println(err)
	}
}

// CreateTemplateCache finds and executes all the templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// ts, err := template.ParseFiles("./templates/"+page, "./templates/base.layout.tmpl")
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
