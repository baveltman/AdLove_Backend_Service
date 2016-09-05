package main

import (
	"googlemaps.github.io/maps"
	"github.com/JustinBeckwith/go-yelp/yelp"
	"github.com/peppage/foursquarego"
)

//PlaceReviewsResult is an aggregation of reviews from multiple sources (Google, Yelp etc)
//Also, provides review analysis insights
type PlaceReviewsResult struct {
	//data about place from Google's Place Details API
	GoogleData maps.PlaceDetailsResult
	//data about place from Yelp's API
	YelpData yelp.Business
	//data about place from Foursquare's API
	FoursquareData foursquarego.Venue
	
}

type FoursquareOauthOptions struct{
	ClientId		string
	ClientSecret	string
}