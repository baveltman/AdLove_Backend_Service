
package main

import (
	"os"
	"errors"
	"encoding/json"
	"net/http"
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
)

func SearchPlace(search string) (maps.AutocompleteResponse, error){
	var client *maps.Client
	var err error
	
	var placesApiKey = os.Getenv(PlacesApiKey)
	if placesApiKey == ""{
		return maps.AutocompleteResponse{}, errors.New("Missing Google Places Api Key")
	}
	
	//TODO: change this to env var
	client, err = maps.NewClient(maps.WithAPIKey(placesApiKey))
	
	if err != nil{
		return maps.AutocompleteResponse{}, errors.New(err.Error())
	}
	
	searchRequest := &maps.PlaceAutocompleteRequest{
		Input:    search,
	}
	
	resp, err := client.PlaceAutocomplete(context.Background(), searchRequest)
	
	if err != nil {
		return maps.AutocompleteResponse{}, errors.New(err.Error())
	}
	
	return resp, nil
}

func PlacesSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	search := r.URL.Query().Get(Search)
	
	if search == "" {
		http.Error(w, "Missing search", http.StatusBadRequest)
		return
	}
	
	if search != "" {
		resp, err := SearchPlace(search)
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