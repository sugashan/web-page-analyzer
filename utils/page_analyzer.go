package utils

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FindHtmlVersion(content string) string {
	if strings.Contains(content, "<!DOCTYPE html>") || strings.Contains(content, "<!docttype html>") {
		return "HTML5"
	} else if strings.Contains(content, "<!DOCTYPE HTML PUBLIC") || strings.Contains(content, "<!docttype html public>") {
		return "HTML 4"
	} else if strings.Contains(content, "<!DOCTYPE XHTML PUBLIC") || strings.Contains(content, "<!docttype xhtml public>") {
		return "XHTML"
	}
	return "Unknown HTML Version."
}

func GetTitle(doc *goquery.Document) string {
	return doc.Find("title").Text()
}

func CountHeadings(doc *goquery.Document) map[string]int {
	headings := make(map[string]int)
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		level := s.Nodes[0].Data
		headings[level]++
	})
	return headings
}

func CountLinks(urlToAnalyze string, doc *goquery.Document) (int, int, int) {
	internalLinks, externalLinks, inaccessibleLinks := 0, 0, 0
	parsedURL, _ := url.Parse(urlToAnalyze)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
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

func HasLoginForm(doc *goquery.Document) bool {
	hasLoginForm := false
	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		if s.Find("input[type=password]").Length() > 0 {
			hasLoginForm = true
		}
	})
	return hasLoginForm
}
