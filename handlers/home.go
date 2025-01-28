// Package handlers for API.
package handlers

import (
	"html/template"
	"log"
	"net/http"
	"webpageanalyzer/models"
)

// HomeHandler handles home requests
func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}
	data := models.PageData{}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
		return
	}
}
