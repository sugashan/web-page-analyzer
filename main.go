// Package main for Application.
package main

import (
	"fmt"
	"log"
	"net/http"
	"webpageanalyzer/handlers"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HomeHandler)

	mux.HandleFunc("POST /api/v1/analysis", handlers.WebPageAnalyzeHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}
