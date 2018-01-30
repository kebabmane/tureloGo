package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// declare DB
var db *gorm.DB

// Category data model
type Category struct {
	gorm.Model
	CategoryName        string
	CategoryImageURL    string
	CategoryDescription string
	FeedsCount          string
	CategoryID          uint
}

// Feed data model
type Feed struct {
	gorm.Model
	FeedName        string
	FeedURL         string
	FeedIcon        string
	FeedsCount      string
	LastFeteched    string
	FeedDescription string
	FeedImageURL    string
	Categories      []Category
}

// Seeding tables:
var categories []Category = []Category{
	Category{CategoryName: "technology", CategoryImageURL: "http://s3.com", CategoryDescription: "this is where we put some technology stuff"},
	Category{CategoryName: "health", CategoryImageURL: "http://s3.com", CategoryDescription: "this is where we put some health stuff"},
	Category{CategoryName: "medical", CategoryImageURL: "http://s3.com", CategoryDescription: "this is where we put some medical stuff"},
}

// Init migrates the database, in the future add a feature flag to know when to migrate
func Init() {

	dbString := os.Getenv("DATABASE_URL")

	fmt.Println("Is this your DB string: ", dbString)
	var err error
	db, err = gorm.Open("postgres", dbString)
	if err != nil {
		panic("Unable to connect to DB")
	}

	db.AutoMigrate(&Category{})
	fmt.Println("We have migrated the database")

	db.Unscoped().Delete(&categories)
	fmt.Println("We have reset the database")

	for _, category := range categories {
		db.Create(&category)
	}
	fmt.Println("We have seeded the categories")
}
