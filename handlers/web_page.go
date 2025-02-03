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

// WebPageAnalyzeHandler handles for Web Page Analyzer Request
func WebPageAnalyzeHandler(w http.ResponseWriter, r *http.Request) {
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

	log.Println("Analyzing : " + url)

	tmpl, err := template.ParseFiles("templates/results.html")
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}

	results, pageErr := service.AnalyzeURL(url)

	log.Println("Completed Analyze. Results.")

	data := models.PageData{
		Results: results,
	}

	if pageErr != nil {
		data.Error = fmt.Sprintf("Error analyzing URL: %v", pageErr)
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
		return
	}
}
