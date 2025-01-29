// Package service for analyzing url opr web page
package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"webpageanalyzer/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

// fetchHTML uses chromedp to get the raw HTML content of the page as pages could contain dynamic parts with javaScript.
func fetchHTML(urlToAnalyze string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var htmlContent string

	err := chromedp.Run(ctx,
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
