package main

import (
	"errors"
	"os"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
	"github.com/gorilla/mux"
	"github.com/JustinBeckwith/go-yelp/yelp"
	"github.com/peppage/foursquarego"
)

const FIVE_MILES_IN_METERS = "8046.72";

func getFoursquareDetails(placeName string, lat float64, lng float64) (foursquarego.Venue, error) {
	var err error
	var api *foursquarego.FoursquareApi
	
	options, err := getFoursquareClientOptions()
	if err != nil {
		return foursquarego.Venue{}, err
	}
	
	api = foursquarego.NewFoursquareApi(options.ClientId, options.ClientSecret)
	
	uv := url.Values{}
	var latStr string = strconv.FormatFloat(lat, 'f', 6, 64)
	var lngStr string = strconv.FormatFloat(lng, 'f', 6, 64)
	
	uv.Set("ll", latStr + "," + lngStr)
	uv.Set("query", placeName)
	uv.Set("radius", FIVE_MILES_IN_METERS)
	
	venues, err := api.Search(uv)
	if err != nil{
		return foursquarego.Venue{}, err
	}
	
	if len(venues) < 1 {
		return foursquarego.Venue{}, nil
	}
	
	venue, err := api.GetVenue(venues[0].ID)
	if err != nil {
		return foursquarego.Venue{}, nil
	}
	
	return venue, nil
	
}

func getYelpDetails(placeName string, location string) (yelp.Business, error){
	var err error
	
	options, err := getYelpClientOptions()
	if err != nil {
		return yelp.Business{}, err
	}
	
	client := yelp.New(options, nil)
	
	resp, err := client.DoSimpleSearch(placeName, location)
	if err != nil {
		return yelp.Business{}, err
	}
	
	if len(resp.Businesses) < 1{
		return yelp.Business{}, nil
	}
	
	var buisId string = resp.Businesses[0].ID
	
	buisData, err := client.GetBusiness(buisId)
	if err != nil {
		return yelp.Business{}, err
	}
	
	return buisData, nil
}

func getGoogleDetails(placeId string) (maps.PlaceDetailsResult, error) {
	var client *maps.Client
	var err error
	
	var placesApiKey = os.Getenv(PlacesApiKey)
	if placesApiKey == ""{
		return maps.PlaceDetailsResult{}, errors.New("Missing Google Places Api Key")
	}
	
	client, err = maps.NewClient(maps.WithAPIKey(placesApiKey))
	
	if err != nil{
		return maps.PlaceDetailsResult{}, err
	}
	
	searchRequest := &maps.PlaceDetailsRequest{
		PlaceID: placeId,
	}
	
	resp, err := client.PlaceDetails(context.Background(), searchRequest)
	
	if err != nil {
		return maps.PlaceDetailsResult{}, err 
	}
	
	return resp, nil
}


func PlacesReviewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	vars := mux.Vars(r)
	
	placeId := vars[PlaceId]
	
	
	if placeId == "" {
		http.Error(w, "Missing placeId", http.StatusBadRequest)
		return
	}
	
	var resp PlaceReviewsResult
	var googleData maps.PlaceDetailsResult
	//get google details
	googleData, err := getGoogleDetails(placeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	
	resp.GoogleData = googleData
	
	//get yelp details
	var yelpData yelp.Business
	yelpData, err = getYelpDetails(googleData.Name, googleData.FormattedAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	resp.YelpData = yelpData
	
	//get Foursquare details
	var foursquareData foursquarego.Venue
	var lat float64 = googleData.Geometry.Location.Lat
	var lng float64 = googleData.Geometry.Location.Lng
	foursquareData, err = getFoursquareDetails(googleData.Name, lat, lng)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	
	resp.FoursquareData = foursquareData

	//return response as json
	js, err := json.Marshal(resp)
	  
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	
}