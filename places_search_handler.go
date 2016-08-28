package main

import (
	"encoding/json"
	"net/http"
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
)

const Search string = "search"

func PlacesSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	search := r.URL.Query().Get("search")
	
	if search == "" {
		http.Error(w, "Missing search", http.StatusBadRequest)
		return
	}
	
	if search != "" {
		var client *maps.Client
		var err error
		
		//TODO: change this to env var
		client, err = maps.NewClient(maps.WithAPIKey("AIzaSyDE-5pwf2NVgurJWZACWe-6JDhS05vjkAo"))
		
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