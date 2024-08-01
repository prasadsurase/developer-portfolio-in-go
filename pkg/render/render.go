package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate is used to render the template
func RenderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	err = parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println(err)
	}
}
