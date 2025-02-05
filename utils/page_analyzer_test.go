package utils

import (
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func stringToReadCloser(s string) io.ReadCloser {
	return io.NopCloser(strings.NewReader(s))
}

func TestFindHTMLVersion(t *testing.T) {
	cases := []struct {
		html     string
		expected string
	}{
		{"<!DOCTYPE html>", "HTML5"},
		{"<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\">", "HTML 4.01"},
		{"<!DOCTYPE XHTML 1.0 Strict PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\">", "XHTML"},
		{"<html></html>", "Unknown"}, // No doctype provided
	}

	for _, tc := range cases {
		t.Run(tc.html, func(t *testing.T) {
			reader := stringToReadCloser(tc.html)
			assert.Equal(t, tc.expected, FindHTMLVersion(reader))
		})
	}
}

func TestGetTitle(t *testing.T) {
	html := `<html><head><title>Test Page</title></head><body></body></html>`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	assert.Equal(t, "Test Page", GetTitle(doc))
}

func TestCountHeadings(t *testing.T) {
	html := `<html><body><h1>Title</h1><h2>Subtitle</h2><h2>Another Subtitle</h2></body></html>`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	expected := map[string]int{"h1": 1, "h2": 2}
	assert.Equal(t, expected, CountHeadings(doc))
}

// func TestCountLinks(t *testing.T) {
// 	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	defer svr.Close()

// 	html := `<html><body>
// 		<a href="` + svr.URL + `">Internal</a>
// 		<a href="https://example.com">External</a>
// 		<a href="invalid">Broken</a>
// 	</body></html>`
// 	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

// 	internal, external, broken := CountLinks(svr.URL, doc)
// 	assert.Equal(t, 1, internal)
// 	assert.Equal(t, 1, external)
// 	assert.Equal(t, 1, broken)
// }

func TestHasLoginForm(t *testing.T) {
	htmlWithLogin := `<html><body><form><input type="password"></form></body></html>`
	htmlWithoutLogin := `<html><body><form><input type="text"></form></body></html>`
	docWithLogin, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlWithLogin))
	docWithoutLogin, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlWithoutLogin))

	assert.True(t, HasLoginForm(docWithLogin))
	assert.False(t, HasLoginForm(docWithoutLogin))
}
