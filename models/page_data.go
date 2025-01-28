// Package models provides data structures.
package models

// PageData represents the data structure used for storing and sharing to templates
type PageData struct {
	Error   string
	Results map[string]interface{}
}
