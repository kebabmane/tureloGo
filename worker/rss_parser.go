package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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
		go crawlFeed(f, ch)
	}

	for i := 0; i < len(feeds); i++ {
		fmt.Println(<-ch)
	}
}

// CrawlFeed allows you to crawl feeds
func crawlFeed(f Feed, ch chan<- string) {
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
		var feedEntry model.FeedEntry
		feedEntry.FeedEntryTitle = i.Title
		feedEntry.FeedEntryAuthor = i.Author.Name
		feedEntry.FeedEntryPublished = i.Published
		feedEntry.FeedEntryLink = i.Link
		feedEntry.FeedEntryContent = i.Content
		feedEntry.FeedID = f.ID
		fmt.Println("this feedEntry: ", feedEntry)
		b, _ := json.Marshal(feedEntry)
		model.CreateFeedEntry(b)
	}
	ch <- "successfully crawled " + f.FeedURL + "\n"
}

func main() {
	// load application configurations in not production
	if os.Getenv("ENV") == "PRODUCTION" {
		fmt.Println("your running in production, did you know that?")
	} else {
		fmt.Println("your running in dev/test, did you know that?")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	model.Init()
	Crawl()
}
