package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// declare DB
var db *gorm.DB

type (
	categoryModel struct {
		gorm.Model
		ID                  int    `json:"id"`
		CategoryName        string `json:"category_name"`
		CategoryImageURL    string `json:"category_image_url"`
		CategoryDescription string `json:"category_description"`
		FeedsCount          string `json:"feeds_count"`
	}
)

func init() {

	// dbString := os.Getenv("DATABASE_URL")

	dbString := "postgres://postgres:postgres@127.0.0.1:5432/turelogo?sslmode=disable"

	fmt.Println("Is this your DB string: ", dbString)
	var err error
	db, err = gorm.Open("postgres", dbString)
	if err != nil {
		panic("Unable to connect to DB")
	}

	db.AutoMigrate(&categoryModel{})
}
