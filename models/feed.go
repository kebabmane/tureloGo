package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

// Feed represents an feed record.
type Feed struct {
	ID                int    `json:"id" db:"id"`
	FeedName          string `json:"feedName" db:"feed_name"`
	FeedURL           string `json:"feedURL" db:"feed_url"`
	NumberFeedEntries int    `json:"numberFeedEntries" db:"number_feed_rntries"`
	FeedIconURL       string `json:"feedIconURL" db:"feed_icon_url"`
	LastFeteched      string `json:"lastFetched" db:"last_fetched"`
	FeedEntriesCount  int    `json:"feedEntriesCount" db:"feed_entries_count"`
	FeedDescription   string `json:"feedDescription" db:"feed_description"`
}

// Validate validates the Feed fields.
func (m Feed) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.FeedName, validation.Required, validation.Length(0, 120)),
	)
}
