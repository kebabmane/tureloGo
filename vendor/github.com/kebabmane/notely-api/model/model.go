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
	CategoryName string `gorm:"unique_index"`
	CategoryID   uint
}

// Label data model
type Label struct {
	gorm.Model
	LabelName string
	LabelID   uint
}

// Todo data model
type Todo struct {
	gorm.Model
	TodoID      uint
	TodoText    string
	TodoDueDate string
	Category    Category
	UserID      uint
	Labels      []Label
}

// User data model
type User struct {
	gorm.Model
	UserID uint
	Todos  []Todo
	Labels []Label
}

// Seeding tables:
var categories []Category = []Category{
	Category{CategoryName: "technology"},
	Category{CategoryName: "health"},
	Category{CategoryName: "medical"},
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
	fmt.Println("We have seeded the database with categories")
}
