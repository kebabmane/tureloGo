package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// declare DB
var db *gorm.DB

// Category data model
type Category struct {
	gorm.Model
	CategoryName        string `gorm:"unique_index"`
	CategoryImageURL    string
	CategoryDescription string
	FeedsCount          string
	CategoryID          uint
}

// Feed data model
type Feed struct {
	gorm.Model
	FeedName        string `gorm:"unique_index"`
	FeedURL         string
	FeedIcon        string
	FeedsCount      string
	LastFeteched    string
	FeedDescription string
	FeedImageURL    string
	FeedLastUpdated time.Time
	Categories      []Category
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

// Init migrates the database, in the future add a feature flag to know when to migrate
func Init() {

	dbString := "postgres://postgres:postgres@127.0.0.1:5432/turelogo?sslmode=disable"

	fmt.Println("Is this your DB string: ", dbString)
	var err error
	db, err = gorm.Open("postgres", dbString)
	if err != nil {
		panic("Unable to connect to DB")
	}

	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Feed{})
	fmt.Println("We have migrated the database")

	db.Unscoped().Delete(&categories)
	db.Unscoped().Delete(&feeds)
	fmt.Println("We have reset the database")

	for _, category := range categories {
		db.Create(&category)
	}
	for _, feed := range feeds {
		db.Create(&feed)
	}
	fmt.Println("We have seeded the database with feeds & categories")
}
