package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate is used to render the template
func RenderTemplate(rw http.ResponseWriter, tmplName string) {
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmplName]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	t.Execute(buf, nil)

	_, err = buf.WriteTo(rw)
	if err != nil {
		fmt.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
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
