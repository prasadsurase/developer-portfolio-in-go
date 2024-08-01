package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const serverPort = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", serverPort)
	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.tmpl")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.tmpl")
}

func renderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println(err)
	}

	err = parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println(err)
	}
}
