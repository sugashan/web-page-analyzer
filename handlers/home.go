package handlers

import (
	"html/template"
	"log"
	"net/http"
	"webpageanalyzer/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}
	data := models.PageData{}
	tmpl.Execute(w, data)
}
