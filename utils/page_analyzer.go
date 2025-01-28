// Package utils for document analyzer.
package utils

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// FindHTMLVersion Finds HTML Version
func FindHTMLVersion(content string) string {
	switch {
	case strings.Contains(content, "<!DOCTYPE html>") || strings.Contains(content, "<!docttype html>"):
		return "HTML 5"
	case strings.Contains(content, "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\">"):
		return "HTML 4.01"
	case strings.Contains(content, "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\">"):
		return "XHTML 1.0 Strict"
	default:
		return "Unknown"
	}

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
