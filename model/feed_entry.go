package model

import (
	"encoding/json"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// FetchAllFeedEntries creates a new feed item and returns the []byte json object and an error.
func FetchAllFeedEntries(id string) ([]byte, error) {

	var feedEntries []FeedEntry

	table := getFeedEntriesTableName()
	feedEntryTable := db.Table(*table)

	err := feedEntryTable.Scan().All(&feedEntries)

	if err != nil {
		log.Println("%+v\n", err)
	}

	if len(feedEntries) <= 0 {
		err := errors.New("Not found")
		{
			return []byte("feed entries not found"), err
		}
	}

	js, err := json.Marshal(feedEntries)
	{
		return js, err
	}
}

// CreateFeedEntry creates a new feed item and returns the []byte json object and an error.
func CreateFeedEntry(b []byte) ([]byte, error) {

	var feedEntry FeedEntry

	err := json.Unmarshal(b, &feedEntry)

	table := getFeedEntriesTableName()
	feedEntryTable := db.Table(*table)

	err = feedEntryTable.Put(&feedEntry).Run()

	if err != nil {
		return []byte("Something went wrong"), err
	}

	return []byte("Feed entry successfully created"), nil
}

// FetchSingleFeedEntry gets a single feed based on param passed, returning []byte and error
func FetchSingleFeedEntry(id string) ([]byte, error) {

	var feedEntry FeedEntry

	table := getFeedEntriesTableName()
	feedEntryTable := db.Table(*table)

	err := feedEntryTable.Get("FeedEntryID", id).One(&feedEntry)

	if feedEntry.FeedEntryID == 0 {
		err := errors.New("Not found")
		return []byte("feed entry not found"), err
	}

	fmt.Println("this is the feed: ", feedEntry)

	js, err := json.Marshal(feedEntry)
	if err != nil {
		js = []byte("Unable to convert feed entry to JSON format")
	}

	return js, err
}

// DeleteFeedEntry deletes the feed entry from the database
func DeleteFeedEntry(id string) ([]byte, error) {

	table := getFeedEntriesTableName()
	feedEntryTable := db.Table(*table)

	var feedEntry FeedEntry
	err := feedEntryTable.Get("FeedEntryID", id).One(&feedEntry)

	if feedEntry.FeedEntryID == 0 {
		err := errors.New("Not found")
		return []byte("feed entry not found"), err
	}

	err = feedEntryTable.Delete("FeedEntryID", id).Run()

	if err != nil {
		log.Println("%+v\n", err)
	}

	return []byte("feed entry deleted"), err
}
