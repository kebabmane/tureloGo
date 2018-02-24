package model

import (
	"encoding/json"
	"errors"
)

// FetchAll is the model function which interfaces with the DB and returns a []byte of the category in json format.
func FetchAllCategories() ([]byte, error) {

	var categories []Category

	db.Find(&categories)

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

// Create creates a new category item and returns the []byte json object and an error.
func CreateCategory(b []byte) ([]byte, error) {

	var category Category

	err := json.Unmarshal(b, &category)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	db.Save(&category)

	return []byte("Category successfully created"), nil
}

// FetchSingle gets a single todo based on param passed, returning []byte and error
func FetchSingleCategory(id string) ([]byte, error) {

	var category Category
	db.First(&category, id)

	if category.ID == 0 {
		err := errors.New("Not found")
		return []byte("category not found"), err
	}

	js, err := json.Marshal(category)
	if err != nil {
		js = []byte("Unable to convert category to JSON format")
	}

	return js, err
}

// Update is the model function for PUT
func UpdateCategory(b []byte, id string) ([]byte, error) {

	var category, updatedCategory Category
	db.First(&category, id)

	if category.ID == 0 {
		err := errors.New("Not found")
		return []byte("category not found"), err
	}

	err := json.Unmarshal(b, &updatedCategory)
	if err != nil {
		return []byte("Malformed input"), err
	}

	db.Model(&category).Update("category_name", updatedCategory.CategoryName)

	js, err := json.Marshal(&category)
	if err != nil {
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// Delete deletes the category from the database
func DeleteCategory(id string) ([]byte, error) {

	var category Category
	db.First(&category, id)

	if category.ID == 0 {
		err := errors.New("Not found")
		return []byte("category not found"), err
	}

	db.Delete(&category)

	js, err := json.Marshal(&category)
	if err != nil {
		panic("Unable to marshal category into json")
	}

	return js, nil
}
