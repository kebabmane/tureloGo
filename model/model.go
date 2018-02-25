package model

import (
	"errors"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

// declare DB
var db *dynamo.DB

// Category data model
type Category struct {
	CategoryName        string `dynamo:"CategoryName"`
	CategoryImageURL    string `dynamo:"DatabaseImageURL"`
	CategoryDescription string `dynamo:"CategoryDescription"`
	FeedsCount          string `dynamo:"FeedsCount"`
	CategoryID          uint   // Hash key
}

// Feed data model
type Feed struct {
	FeedID          uint
	FeedName        string `dynamo:"FeedName"`
	FeedURL         string `dynamo:"FeedURL"`
	FeedIcon        string `dynamo:"FeedIcon"`
	FeedsCount      string `dynamo:"FeedsCount"`
	LastFeteched    string `dynamo:"LastFetched"`
	FeedDescription string `dynamo:"FeedDescriptiom"`
	FeedImageURL    string `dynamo:"FeedImageURL"`
	FeedLastUpdated time.Time
	Categories      []Category
	FeedEntry       []FeedEntry
}

// FeedEntry data model
type FeedEntry struct {
	FeedEntryTitle            string `dynamo:"FeedEntryTitle"`
	FeedEntryURL              string `dynamo:"FeedEntryURL"`
	FeedEntryPublished        string `dynamo:"FeedEntryPublished"`
	FeedEntryAuthor           string `dynamo:"FeedEntryAuthor"`
	FeedEntryContent          string `dynamo:"FeedEntryContent"`
	FeedEntryContentSanitized string `dynamo:"FeedEntryContentSanitized"`
	FeedEntryLink             string `dynamo:"FeedEntryLink"`
	FeedID                    uint
}

// Seeding tables:
var categories []Category = []Category{
	Category{CategoryName: "technology", CategoryImageURL: "https://www.imore.com/sites/imore.com/files/styles/xlarge/public/field/image/2016/03/ipad-mini-ipad-air-ipad-pro-stack-snow-hero.jpg?itok=ir4jkST2", CategoryDescription: "this is where we put some technology stuff"},
	Category{CategoryName: "health", CategoryImageURL: "https://lorempixel.com/600/300/food/5/", CategoryDescription: "this is where we put some health stuff"},
	Category{CategoryName: "medical", CategoryImageURL: "https://lorempixel.com/600/300/food/5/", CategoryDescription: "this is where we put some medical stuff"},
}

var feeds []Feed = []Feed{
	Feed{FeedName: "The Verge -  All Posts", FeedURL: "http://theverge.com/rss/index.xml", FeedDescription: "this is where we put some technology stuff", FeedIcon: "https://cdn.vox-cdn.com/community_logos/52801/VER_Logomark_32x32..png"},
}

// ErrorBadRequest is the bad request string
var ErrorBadRequest = errors.New("Bad request")

// ErrorInternalServer is the internal server error
var ErrorInternalServer = errors.New("Something went wrong")

// ErrorForbidden is the forbidden error
var ErrorForbidden = errors.New("Forbidden")

// ErrorNotFound is the not found error
var ErrorNotFound = errors.New("Not found")

// Init migrates the database, in the future add a feature flag to know when to migrate
func Init() {

	feedsTable := db.Table(getFeedsTableName)
	err := feedsTable.Put(feeds).Run()
	if err != nil {
		log.Println("%+v\n", err)
	}

}

func getFeedsTableName() string {
	// Setup the table names as required for models
	var tableName = aws.String(os.Getenv("DATABASE_FEEDS_TABLE"))
	// return the table name as a string
	return tableName
}

func getCategoriesTableName() string {
	// Setup the table names as required for models
	var tableName = aws.String(os.Getenv("DATABASE_CATEGORIES_TABLE"))
	// return the table name as a string
	return tableName
}

func getFeedEntriesTableName() string {
	// Setup the table names as required for models
	var tableName = aws.String(os.Getenv("DATABASE_FEEDENTRIES_TABLE"))
	// return the table name as a string
	return tableName
}
