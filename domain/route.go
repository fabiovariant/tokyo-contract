package domain

import (
	"net/http"
)

// Route is a type that simplify the creatation of routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

// Routes is Array of routes
type Routes []Route
