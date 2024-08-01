package render

import (
	"fmt"
	"html/template"
	"net/http"
)

var templatesCache = make(map[string]*template.Template)

// RenderTemplate is used to render the template
func RenderTemplate(rw http.ResponseWriter, tmplName string) {
	var tmpl *template.Template
	var err error

	_, inMap := templatesCache[tmplName]
	if !inMap {
		fmt.Println("creating template and adding to cache")
		err = createTemplateCache(tmplName)
		if err != nil {
			fmt.Printf("error creating template cache: %v\n", err)
		}
	} else {
		fmt.Printf("using cached template: %v\n", tmplName)
	}

	tmpl = templatesCache[tmplName]
	err = tmpl.Execute(rw, nil)
	if err != nil {
		fmt.Printf("error executing template: %v\n", err)
	}
}

// createTemplateCache creates the cache for templates using map
func createTemplateCache(tmplName string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", tmplName),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	templatesCache[tmplName] = tmpl
	return nil
}
