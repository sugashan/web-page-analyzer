// Package service for analyzing url opr web page
package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"webpageanalyzer/config"
	"webpageanalyzer/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

// fetchHTML uses chromedp to get the raw HTML content of the page as pages could contain dynamic parts with javaScript.
func fetchHTML(urlToAnalyze string) (string, error) {
	resp, err := http.Get(urlToAnalyze)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-2xx responses
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf(" HTTP %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	timeout := config.GetRequestTimeout()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var htmlContent string

	err = chromedp.Run(ctx,
		chromedp.Navigate(urlToAnalyze),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		return "", err
	}

	return htmlContent, nil
}

// AnalyzeURL does web scraping for given URL
func AnalyzeURL(urlToAnalyze string) (map[string]interface{}, error) {

	htmlContent, err := fetchHTML(urlToAnalyze)

	if err != nil {
		return nil, fmt.Errorf("error fetching the URL: %v", err)
	}

	log.Println("Found the Page")

	htmlVersion := utils.FindHTMLVersion(htmlContent)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
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
