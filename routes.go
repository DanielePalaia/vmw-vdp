package main

import (
	"net/http"
	"vmw-vdp/controllers"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Route struct contains route configuration
type Route struct {
	name    string
	method  string
	pattern string
	handler http.Handler
}

// Routes ia a vector of Route
type Routes []Route

// NewRouter returns a router to manage different routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.handler)
	}

	return router
}

var routes = Routes{
	// Returns all elements of the collection todos
	Route{
		"getMetrics",
		"GET",
		"/service",
		http.HandlerFunc(controllers.HandleServiceRequest),
	},
	Route{
		"getMetrics",
		"GET",
		"/metrics",
		promhttp.Handler(),
	},
}
