package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var loginController LoginController

var routes = Routes{
	Route{
		"index",
		"/",
		"GET",
		loginController.Index,
	},
	Route{
		"indexPost",
		"/",
		"POST",
		loginController.Index,
	},
}
