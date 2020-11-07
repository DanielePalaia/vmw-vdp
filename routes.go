package main

import (
	"net/http"
	"vmw-vdp/controllers"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}

var routes = Routes{
	// Returns all elements of the collection todos
	Route{
		"getMetrics",
		"GET",
		"/service",
		http.HandlerFunc(controllers.HandleMetricsRequest),
	},
	Route{
		"getMetrics",
		"GET",
		"/metrics",
		promhttp.Handler(),
	},
}
