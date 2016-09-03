package main

import (
	"errors"
	"os"
	"encoding/json"
	"net/http"
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
	"github.com/gorilla/mux"
)

func getGoogleDetails(placeId string) (maps.PlaceDetailsResult, error) {
	var client *maps.Client
	var err error
	
	var placesApiKey = os.Getenv(PlacesApiKey)
	if placesApiKey == ""{
		return maps.PlaceDetailsResult{}, errors.New("Missing Google Places Api Key")
	}
	
	//TODO: change this to env var
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
	
	googleData, err := getGoogleDetails(placeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	
	resp.GoogleData = googleData

	//return response as json
	js, err := json.Marshal(resp)
	  
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	
}