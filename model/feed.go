package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FetchAllFeeds is the model function which interfaces with the DB and returns a []byte of the category in json format.
func FetchAllFeeds() ([]byte, error) {

	var feeds []Feed

	db.Find(&feeds)

	if len(feeds) <= 0 {
		err := errors.New("Not found")
		{
			return []byte("feeds not found"), err
		}
	}

	js, err := json.Marshal(feeds)
	{
		return js, err
	}
}

// CreateFeed creates a new category item and returns the []byte json object and an error.
func CreateFeed(b []byte) ([]byte, error) {

	var feed Feed

	err := json.Unmarshal(b, &feed)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	db.Save(&feed)

	return []byte("Feed successfully created"), nil
}

// FetchSingleFeed gets a single feed based on param passed, returning []byte and error
func FetchSingleFeed(id string) ([]byte, error) {

	var feed Feed
	db.First(&feed, id)

	if feed.ID == 0 {
		err := errors.New("Not found")
		return []byte("feed not found"), err
	}

	fmt.Println("this is the feed: ", feed)

	js, err := json.Marshal(feed)
	if err != nil {
		js = []byte("Unable to convert feed to JSON format")
	}

	return js, err
}

// UpdateFeed is the model function for PUT
func UpdateFeed(b []byte, id string) ([]byte, error) {

	var feed, updatedFeed Feed
	db.First(&feed, id)

	if feed.ID == 0 {
		err := errors.New("Not found")
		return []byte("feed not found"), err
	}

	err := json.Unmarshal(b, &updatedFeed)
	if err != nil {
		return []byte("Malformed input"), err
	}

	db.Model(&feed).Update("feed_name", updatedFeed.FeedName)
	db.Model(&feed).Update("feed_url", updatedFeed.FeedURL)
	db.Model(&feed).Update("feed_icon", updatedFeed.FeedIcon)
	db.Model(&feed).Update("feeds_count", updatedFeed.FeedsCount)
	db.Model(&feed).Update("last_fetched", updatedFeed.LastFeteched)
	db.Model(&feed).Update("feed_description", updatedFeed.FeedName)
	db.Model(&feed).Update("Feed_image_url", updatedFeed.FeedImageURL)

	js, err := json.Marshal(&feed)
	if err != nil {
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// DeleteFeed deletes the feed from the database
func DeleteFeed(id string) ([]byte, error) {

	var feed Feed
	db.First(&feed, id)

	if feed.ID == 0 {
		// w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte("Todo not found"))
		// return
	}

	db.Delete(&feed)

	js, err := json.Marshal(&feed)
	if err != nil {
		panic("Unable to marshal feed into json")
	}

	return js, nil
}
