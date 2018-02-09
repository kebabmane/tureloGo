package model

import (
	"encoding/json"
	"errors"
)

// FetchAllFeedEntries is the model function which interfaces with the DB and returns a []byte of the category in json format.
func FetchAllFeedEntries(id string) ([]byte, error) {

	var feedEntries []FeedEntry

	db.Where("feed_id = ?", id).Find(&feedEntries)

	if len(feedEntries) <= 0 {
		err := errors.New("Not found")
		{
			return []byte("feedEntries not found"), err
		}
	}

	js, err := json.Marshal(feedEntries)
	{
		return js, err
	}
}

// CreateFeedEntry creates a new category item and returns the []byte json object and an error.
func CreateFeedEntry(b []byte) ([]byte, error) {

	var feedEntry FeedEntry

	err := json.Unmarshal(b, &feedEntry)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	db.Save(&feedEntry)

	return []byte("FeedEntry successfully created"), nil
}

// FetchSingleFeedEntry gets a single feed based on param passed, returning []byte and error
func FetchSingleFeedEntry(id string) ([]byte, error) {

	var feedEntry FeedEntry
	db.First(&feedEntry, id)

	if feedEntry.ID == 0 {
		err := errors.New("Not found")
		return []byte("feedEntry not found"), err
	}

	js, err := json.Marshal(feedEntry)
	if err != nil {
		js = []byte("Unable to convert feedEntry to JSON format")
	}

	return js, err
}

// UpdateFeedEntry is the model function for PUT
func UpdateFeedEntry(b []byte, id string) ([]byte, error) {

	var feedEntry, updatedFeedEntry FeedEntry
	db.First(&feedEntry, id)

	if feedEntry.ID == 0 {
		err := errors.New("Not found")
		return []byte("feedEntry not found"), err
	}

	err := json.Unmarshal(b, &updatedFeedEntry)
	if err != nil {
		return []byte("Malformed input"), err
	}

	js, err := json.Marshal(&feedEntry)
	if err != nil {
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// DeleteFeedEntry deletes the feed from the database
func DeleteFeedEntry(id string) ([]byte, error) {

	var feedEntry FeedEntry
	db.First(&feedEntry, id)

	if feedEntry.ID == 0 {
		// w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte("Todo not found"))
		// return
	}

	db.Delete(&feedEntry)

	js, err := json.Marshal(&feedEntry)
	if err != nil {
		panic("Unable to marshal feed into json")
	}

	return js, nil
}
