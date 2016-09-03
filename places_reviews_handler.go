package main

import (
	"os"
	"encoding/json"
	"net/http"
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
	"github.com/gorilla/mux"
)


func PlacesReviewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	vars := mux.Vars(r)
	
	placeId := vars[PlaceId]
	
	
	if placeId == "" {
		http.Error(w, "Missing placeId", http.StatusBadRequest)
		return
	}
	
	if placeId != "" {
		var client *maps.Client
		var err error
		
		var placesApiKey = os.Getenv(PlacesApiKey)
		if placesApiKey == ""{
			http.Error(w, "Missing Google Places Api Key", http.StatusInternalServerError)
		    return
		}
		
		//TODO: change this to env var
		client, err = maps.NewClient(maps.WithAPIKey(placesApiKey))
		
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		    return
		}
		
		searchRequest := &maps.PlaceDetailsRequest{
			PlaceID: placeId,
		}
		
		resp, err := client.PlaceDetails(context.Background(), searchRequest)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		    return
		}
		
		js, err := json.Marshal(resp)
	  
	  if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	  }
	  
	  w.WriteHeader(http.StatusOK)
	  w.Write(js)

	}
	
	
}