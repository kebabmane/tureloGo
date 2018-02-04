package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kebabmane/tureloGo/model"
	"github.com/mmcdole/gofeed"
)

// declare DB
var db *gorm.DB

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
}

// Crawl gets the list of feeds, then pumps each feed through to the CrawlFeed function
func Crawl() {

	ch := make(chan string)

	feeds, err := model.FetchAllFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal string into structs.
	var feed []Feed
	json.Unmarshal(feeds, &feed)

	for _, f := range feed {
		go CrawlFeed(f, ch)
	}

	for i := 0; i < len(feeds); i++ {
		fmt.Println(<-ch)
	}
}

// CrawlFeed allows you to crawl feeds
func CrawlFeed(f Feed, ch chan<- string) {
	c := &http.Client{
		// give up after 5 seconds
		Timeout: 5 * time.Second,
	}

	fp := gofeed.NewParser()
	fp.Client = c

	feed, err := fp.ParseURL(f.FeedURL)
	if err != nil {
		fmt.Println(err)
		ch <- "failed to fetch and parse for " + f.FeedURL + "\n"
		return
	}

	for _, i := range feed.Items {
		fmt.Printf("storing item: %s", i.Link)
		fmt.Printf("the item: ", i)
	}
	ch <- "successfully crawled " + f.FeedURL + "\n"
}

func main() {
	model.Init()
	Crawl()
}
