// Package handlers for WebPage Analyzer API.
package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"webpageanalyzer/models"
	"webpageanalyzer/service"
)

// WebPageeHandler handles for Web Page Analyzer Request
func WebPageeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received to analyze")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	log.Println("Analyzing :" + url)

	results, err := service.AnalyzeURL(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error analyzing URL: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/results.html")
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}

	log.Println("Completed Analyze. Results.")

	err = tmpl.Execute(w, models.PageData{
		Results: results,
	})

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
		return
	}
}
