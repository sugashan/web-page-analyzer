package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"webpageanalyzer/models"
	"webpageanalyzer/utils"

	"github.com/PuerkitoBio/goquery"
)

func analyzeURL(urlToAnalyze string) (map[string]interface{}, error) {
	resp, err := http.Get(urlToAnalyze)
	if err != nil {
		return nil, fmt.Errorf("error fetching the URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to reach page, status code: %d", resp.StatusCode)
	}

	buffer := make([]byte, 1024)
	_, err = resp.Body.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	log.Println("Found the Page")

	content := string(buffer)
	htmlVersion := utils.FindHtmlVersion(content)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
	}

	title := utils.GetTitle(doc)

	headings := utils.CountHeadings(doc)

	internalLinks, externalLinks, inaccessibleLinks := utils.CountLinks(urlToAnalyze, doc)

	hasLoginForm := utils.HasLoginForm(doc)

	results := map[string]interface{}{
		"URL":               urlToAnalyze,
		"HTMLVersion":       htmlVersion,
		"Title":             title,
		"Headings":          headings,
		"InternalLinks":     internalLinks,
		"ExternalLinks":     externalLinks,
		"InaccessibleLinks": inaccessibleLinks,
		"HasLoginForm":      hasLoginForm,
	}
	return results, nil
}

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

	results, err := analyzeURL(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error analyzing URL: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/results.html")
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}

	log.Println("Completed Analyze. Results.")

	tmpl.Execute(w, models.PageData{
		Results: results,
	})
}
