package model

import (
	"encoding/json"

	"github.com/kebabmane/tureloGo/config"
)

// FetchAllFeeds is the model function which interfaces with the DB and returns a []byte of the feed in json format.
func FetchAllFeeds() ([]byte, error) {

	var feeds []Feed
	db := config.GetDB()

	db.Find(&feeds)

	js, err := json.Marshal(feeds)
	{
		return js, err
	}
}

// CreateFeed creates a new feed item and returns the []byte json object and an error.
func CreateFeed(b []byte) ([]byte, error) {

	return nil, nil
}

// FetchSingleFeed gets a single feed based on param passed, returning []byte and error
func FetchSingleFeed(id string) ([]byte, error) {

	var feed Feed
	db := config.GetDB()

	db.First(&feed, id)

	js, err := json.Marshal(feed)
	if err != nil {
		js = []byte("Unable to convert feed to JSON format")
	}

	return js, err
}

// UpdateFeed is the model function for PUT
func UpdateFeed(b []byte, id string) ([]byte, error) {

	return nil, nil
}

// DeleteFeed deletes the feed from the database
func DeleteFeed(id string) ([]byte, error) {

	return nil, nil
}
