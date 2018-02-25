package model

import (
	"encoding/json"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// FetchAllFeeds is the model function which interfaces with the DB and returns a []byte of the feed in json format.
func FetchAllFeeds() ([]byte, error) {

	var feeds []Feed

	table := getFeedsTableName()
	feedTable := db.Table(*table)

	err := feedTable.Scan().All(&feeds)

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

// CreateFeed creates a new feed item and returns the []byte json object and an error.
func CreateFeed(b []byte) ([]byte, error) {

	var feed Feed

	err := json.Unmarshal(b, &feed)

	table := getFeedsTableName()
	feedTable := db.Table(*table)

	err = feedTable.Put(&feed).Run()

	if err != nil {
		return []byte("Something went wrong"), err
	}

	return []byte("Feed successfully created"), nil
}

// FetchSingleFeed gets a single feed based on param passed, returning []byte and error
func FetchSingleFeed(id string) ([]byte, error) {

	var feed Feed

	table := getFeedsTableName()
	feedTable := db.Table(*table)

	err := feedTable.Get("FeedID", id).One(&feed)

	if feed.FeedID == 0 {
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

	table := getFeedsTableName()
	feedTable := db.Table(*table)

	var feed, updatedFeed Feed

	if feed.FeedID == 0 {
		err := errors.New("Not found")
		return []byte("feed not found"), err
	}

	err := json.Unmarshal(b, &updatedFeed)
	if err != nil {
		return []byte("Malformed input"), err
	}

	feedTable.Update("feed_name", updatedFeed.FeedName)
	feedTable.Update("feed_url", updatedFeed.FeedURL)
	feedTable.Update("feed_icon", updatedFeed.FeedIcon)
	feedTable.Update("feeds_count", updatedFeed.FeedsCount)
	feedTable.Update("last_fetched", updatedFeed.LastFeteched)
	feedTable.Update("feed_description", updatedFeed.FeedName)
	feedTable.Update("Feed_image_url", updatedFeed.FeedImageURL)

	js, err := json.Marshal(&feed)
	if err != nil {
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// DeleteFeed deletes the feed from the database
func DeleteFeed(id string) ([]byte, error) {

	table := getFeedsTableName()
	feedTable := db.Table(*table)

	var feed Feed
	err := feedTable.Get("FeedID", id).One(&feed)

	if feed.FeedID == 0 {
		err := errors.New("Not found")
		return []byte("feed not found"), err
	}

	err = feedTable.Delete("FeedID", id).Run()

	if err != nil {
		log.Println("%+v\n", err)
	}

	return []byte("feed deleted"), err
}
