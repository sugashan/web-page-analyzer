// Package main for Application.
package main

import (
	"fmt"
	"log"
	"net/http"
	"webpageanalyzer/config"
	"webpageanalyzer/handlers"
)

func main() {
	err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HomeHandler)

	mux.HandleFunc("POST /api/v1/analysis", handlers.WebPageAnalyzeHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}
