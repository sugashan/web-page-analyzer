// Package main for Application.
package main

import (
	"fmt"
	"log"
	"net/http"
	"webpageanalyzer/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/analyze", handlers.WebPageeHandler)

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
