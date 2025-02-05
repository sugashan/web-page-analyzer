// Package utils for document analyzer.
package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// FindHTMLVersion Finds HTML Version
func FindHTMLVersion(body io.ReadCloser) string {
	defer body.Close()
	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return "Unknown"
		case html.DoctypeToken:
			token := tokenizer.Token()
			doctype := strings.ToLower(token.Data)

			if strings.Contains(doctype, "html") {
				if doctype == "html" {
					return "HTML5"
				} else if strings.Contains(doctype, "xhtml") {
					return "XHTML"
				} else if strings.Contains(doctype, "4.01") {
					return "HTML 4.01"
				} else if strings.Contains(doctype, "3.2") {
					return "HTML 3.2"
				}
			}
			return "Unknown Doctype: " + doctype
		}
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
	internalLinks, externalLinks := 0, 0

	inaccessibleLinks := make(chan int, 100)
	var wg sync.WaitGroup

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

		// parellal link access check.
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			if !isLinkAccessible(link) {
				inaccessibleLinks <- 1
			}
		}(href)
	})

	// Wait for all goroutines to finish
	wg.Wait()
	close(inaccessibleLinks)

	// Count inaccessible links from channel
	inaccessibleCount := 0
	for range inaccessibleLinks {
		inaccessibleCount++
	}

	return internalLinks, externalLinks, inaccessibleCount
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

// isLinkAccessible checks if a link is accessible
func isLinkAccessible(link string) bool {
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()
	return true
}
