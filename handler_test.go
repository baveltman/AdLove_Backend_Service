package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"googlemaps.github.io/maps"
)

var placePrediction maps.AutocompletePrediction
var place maps.PlaceDetailsResult

func TestSearchPlace(t *testing.T) {
	var query = "Bedlam Seattle"
	resp, err := SearchPlace(query)
	assert.Nil(t, err, "err is not nil")
	assert.NotNil(t, resp, "resp is nil")
	assert.True(t, len(resp.Predictions) > 0, "No predictions retrieved")
	placePrediction = resp.Predictions[0]
}

func TestGetGoogleDetails(t *testing.T) {
	var placeId = placePrediction.PlaceID
	resp, err := getGoogleDetails(placeId)
	assert.Nil(t, err, "err is not nil")
	assert.NotNil(t, resp, "resp is nil")
	place = resp
}

func TestGetYelpDetails(t *testing.T) {
	var placeName = place.Name
	var placeAddress = place.FormattedAddress
	resp, err := getYelpDetails(placeName, placeAddress)
	assert.Nil(t, err, "err is not nil")
	assert.NotNil(t, resp, "resp is nil")
	assert.True(t, resp.ID != "", "yelp Business has no ID")
}

func TestGetFoursquareDetails(t *testing.T) {
	var query = place.Name
	var lat = place.Geometry.Location.Lat
	var lng = place.Geometry.Location.Lng
	resp, err := getFoursquareDetails(query, lat, lng)
	assert.Nil(t, err, "err is not nil")
	assert.NotNil(t, resp, "resp is nil")
	assert.True(t, resp.ID != "", "Foursquare venue has no ID")
}