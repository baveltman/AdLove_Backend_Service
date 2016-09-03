package main

import (
	"googlemaps.github.io/maps"
)

//PlaceReviewsResult is an aggregation of reviews from multiple sources (Google, Yelp etc)
//Also, provides review analysis insights
type PlaceReviewsResult struct {
	//meta data about place from Google's Place Details API
	GoogleData maps.PlaceDetailsResult
	
}