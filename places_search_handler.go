package main

import (
	"os"
	"encoding/json"
	"net/http"
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
)

const Search string = "search"
const PlacesApiKey = "PLACES_API_KEY"

func PlacesSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	search := r.URL.Query().Get(Search)
	
	if search == "" {
		http.Error(w, "Missing search", http.StatusBadRequest)
		return
	}
	
	if search != "" {
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
		
		searchRequest := &maps.PlaceAutocompleteRequest{
			Input:    search,
		}
		
		resp, err := client.PlaceAutocomplete(context.Background(), searchRequest)
		
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