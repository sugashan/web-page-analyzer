// Package utils for document analyzer.
package utils

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// FindHTMLVersion Finds HTML Version
func FindHTMLVersion(htmlContent string) string {
	switch {
	case strings.Contains(htmlContent, "html"):
		return "HTML5"
	case strings.Contains(htmlContent, "xhtml 1.0"):
		return "XHTML 1.0"
	case strings.Contains(htmlContent, "html 4.01"):
		return "HTML 4.01"
	case strings.Contains(htmlContent, "html 3.2"):
		return "HTML 3.2"
	case strings.Contains(htmlContent, "html 2.0"):
		return "HTML 2.0"
	}
	return "Unknown"
}

// GetTitle Finds Document Title
func GetTitle(doc *goquery.Document) string {
	return doc.Find("title").Text()
}

// CountHeadings Counts Headings in Document
func CountHeadings(doc *goquery.Document) map[string]int {
	headings := make(map[string]int)
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(_ int, s *goquery.Selection) {
		level := s.Nodes[0].Data
		headings[level]++
	})
	return headings
}

// CountLinks Counts Links in Document
func CountLinks(urlToAnalyze string, doc *goquery.Document) (int, int, int) {
	internalLinks, externalLinks, inaccessibleLinks := 0, 0, 0
	parsedURL, _ := url.Parse(urlToAnalyze)

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		linkURL, err := url.Parse(href)
		if err != nil {
			return
		}

		if linkURL.Host == parsedURL.Host || linkURL.Host == "" {
			internalLinks++
		} else {
			externalLinks++
		}

		resp, err := http.Get(href)
		if err != nil || resp.StatusCode != http.StatusOK {
			inaccessibleLinks++
		}
	})
	return internalLinks, externalLinks, inaccessibleLinks
}

// HasLoginForm Finds Login Form
func HasLoginForm(doc *goquery.Document) bool {
	hasLoginForm := false
	doc.Find("form").Each(func(_ int, s *goquery.Selection) {
		if s.Find("input[type=password]").Length() > 0 {
			hasLoginForm = true
		}
	})
	return hasLoginForm
}
