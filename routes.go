package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"IndexHandler",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"PlacesReviews",
		"GET",
		"/reviews/{placeId}",
		PlacesReviewsHandler,
	},
	Route{
		"PlacesSearch",
		"GET",
		"/places",
		PlacesSearchHandler,
	},
	
}