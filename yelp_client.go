package main

import (
	"errors"
	"os"
	"github.com/JustinBeckwith/go-yelp/yelp"
)

func getClientOptions() (*yelp.AuthOptions, error) {
	var o *yelp.AuthOptions
	
	o = &yelp.AuthOptions{
		ConsumerKey:       os.Getenv(YelpConsumerKey),
		ConsumerSecret:    os.Getenv(YelpConsumerSecret),
		AccessToken:       os.Getenv(YelpAccessToken),
		AccessTokenSecret: os.Getenv(YelpAccessTokenSecret),
	}
	
	if o.ConsumerKey == "" {
		return o, errors.New("Missing Yelp Consumer Key")
	}
	
	if o.ConsumerSecret == "" {
		return o, errors.New("Missing Yelp Consumer Secret")
	}
	
	if o.AccessToken == "" {
		return o, errors.New("Missing Yelp Access Token")
	}
	
	if o.AccessTokenSecret == "" {
		return o, errors.New("Missing Yelp Access Token Secret")
	}
	
	return o, nil
}