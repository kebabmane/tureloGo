package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Category data model
type Category struct {
	gorm.Model
	CategoryName        string
	CategoryImageURL    string
	CategoryDescription string
	FeedsCount          string
	CategoryID          uint // Hash key
}

// Feed data model
type Feed struct {
	gorm.Model
	FeedID          uint
	FeedName        string
	FeedURL         string
	FeedIcon        string
	FeedsCount      string
	LastFeteched    string
	FeedDescription string
	FeedImageURL    string
	FeedLastUpdated time.Time
	Categories      []Category
	FeedEntry       []FeedEntry
}

// FeedEntry data model
type FeedEntry struct {
	gorm.Model
	FeedEntryID               uint
	FeedEntryTitle            string
	FeedEntryURL              string
	FeedEntryPublished        string
	FeedEntryAuthor           string
	FeedEntryContent          string
	FeedEntryContentSanitized string
	FeedEntryLink             string
	FeedID                    uint
}

// ErrorBadRequest is the bad request string
var ErrorBadRequest = errors.New("Bad request")

// ErrorInternalServer is the internal server error
var ErrorInternalServer = errors.New("Something went wrong")

// ErrorForbidden is the forbidden error
var ErrorForbidden = errors.New("Forbidden")

// ErrorNotFound is the not found error
var ErrorNotFound = errors.New("Not found")
