package models

import "github.com/go-ozzo/ozzo-validation"

// Artist represents an artist record.
type Feed struct {
	Id       int    `json:"id" db:"id"`
	FeedName string `json:"feedName" db:"feedName"`
}

// Validate validates the Artist fields.
func (m Feed) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.FeedName, validation.Required, validation.Length(0, 120)),
	)
}
