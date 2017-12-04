package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

// Feed represents an feed record.
type FeedEntry struct {
	ID            int    `json:"id" db:"id"`
	FeedEntryName string `json:"feedEntryName" db:"feed_entry_name"`
}

// Validate validates the Feed fields.
func (m FeedEntry) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.FeedEntryName, validation.Required, validation.Length(0, 120)),
	)
}
