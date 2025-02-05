// Package service for analyzing url opr web page
package service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"webpageanalyzer/utils"

	"github.com/PuerkitoBio/goquery"
)

// fetchHTML uses chromedp to get the raw HTML content of the page as pages could contain dynamic parts with javaScript.
func fetchHTML(urlToAnalyze string) (*goquery.Document, string, error) {
	resp, err := http.Get(urlToAnalyze)
	if err != nil {
		return nil, "Unknown", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-2xx responses
	if resp.StatusCode >= 400 {
		return nil, "Unknown", fmt.Errorf(" HTTP %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	log.Println("Found the Page")

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "Unknown", fmt.Errorf("failed to read HTTP response: %w", err)
	}
	bodyCopy1 := io.NopCloser(bytes.NewReader(bodyBytes)) // For tokenizer
	bodyCopy2 := io.NopCloser(bytes.NewReader(bodyBytes)) // For goquery

	htmlVersion := utils.FindHTMLVersion(bodyCopy1)

	doc, err := goquery.NewDocumentFromReader(bodyCopy2)
	if err != nil {
		return nil, htmlVersion, fmt.Errorf("error parsing HTML: %v", err)
	}

	return doc, htmlVersion, nil
}

// AnalyzeURL does web scraping for given URL
func AnalyzeURL(urlToAnalyze string) (map[string]interface{}, error) {

	doc, htmlVersion, err := fetchHTML(urlToAnalyze)

	if err != nil {
		return nil, fmt.Errorf("error fetching the URL: %v", err)
	}

	title := utils.GetTitle(doc)
	log.Println("Title extracted.")

	log.Println("Counting the headings.")
	headings := utils.CountHeadings(doc)

	internalLinks, externalLinks, inaccessibleLinks := utils.CountLinks(urlToAnalyze, doc)
	log.Println("Counting link Completed")

	log.Println("Looking for login page.")
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
