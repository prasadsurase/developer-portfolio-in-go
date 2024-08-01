package main

import (
	"fmt"
	"log"
	"net/http"
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
