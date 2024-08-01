package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prasadsurase/developer-portfolio-in-go/pkg/handlers"
)

const serverPort = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", serverPort)
	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
