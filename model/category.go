package model

import (
	"encoding/json"
	"errors"

	log "github.com/sirupsen/logrus"
)

// FetchAllCategories is the model function which interfaces with the DB and returns a []byte of the category in json format.
func FetchAllCategories() ([]byte, error) {

	var categories []Category

	table := getCategoriesTableName()
	categoryTable := db.Table(table)

	err := categoryTable.Scan().All(&categories)

	if err != nil {
		log.Println("%+v\n", err)
	}

	if len(categories) <= 0 {
		err := errors.New("Not found")
		{
			return []byte("categories not found"), err
		}
	}

	js, err := json.Marshal(categories)
	{
		return js, err
	}
}

// CreateCategory Create creates a new category item and returns the []byte json object and an error.
func CreateCategory(b []byte) ([]byte, error) {

	var category Category

	err := json.Unmarshal(b, &category)

	table := getCategoriesTableName()
	categoryTable := db.Table(table)

	err = categoryTable.Put(&category).Run()

	if err != nil {
		log.Println("%+v\n", err)
		return []byte("Something went wrong"), err
	}

	return []byte("Category successfully created"), nil
}

// FetchSingleCategory gets a single todo based on param passed, returning []byte and error
func FetchSingleCategory(id string) ([]byte, error) {

	table := getCategoriesTableName()
	categoryTable := db.Table(table)

	var category Category

	err := categoryTable.Get("CategoryID", id).One(&category)

	if category.CategoryID == 0 {
		err := errors.New("Not found")
		return []byte("category not found"), err
	}

	js, err := json.Marshal(category)
	if err != nil {
		log.Println("%+v\n", err)
		js = []byte("Unable to convert category to JSON format")
	}

	return js, err
}

// UpdateCategory is the model function for PUT
func UpdateCategory(b []byte, id string) ([]byte, error) {

	table := getCategoriesTableName()
	categoryTable := db.Table(table)

	var category, updatedCategory Category

	if category.CategoryID == 0 {
		err := errors.New("Not found")
		return []byte("category not found"), err
	}

	err := json.Unmarshal(b, &updatedCategory)
	if err != nil {
		log.Println("%+v\n", err)
		return []byte("Malformed input"), err
	}

	categoryTable.Update("category_name", updatedCategory.CategoryName)
	categoryTable.Update("category_image_url", updatedCategory.CategoryImageURL)
	categoryTable.Update("category_description", updatedCategory.CategoryDescription)
	categoryTable.Update("feeds_count", updatedCategory.FeedsCount)

	// get the current category
	err = categoryTable.Get("CategoryID", id).One(&category)

	js, err := json.Marshal(&category)
	if err != nil {
		log.Println("%+v\n", err)
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// DeleteCategory deletes the categoryo from the database
func DeleteCategory(id string) ([]byte, error) {

	table := getCategoriesTableName()
	categoryTable := db.Table(table)

	var category Category
	err := categoryTable.Get("CategoryID", id).One(&category)

	if category.CategoryID == 0 {
		err := errors.New("Not found")
		return []byte("category not found"), err
	}

	err = categoryTable.Delete("CategoryID", id).Run()

	if err != nil {
		log.Println("%+v\n", err)
	}

	return []byte("category deleted"), err
}
