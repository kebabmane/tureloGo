package model

import (
	"encoding/json"
	"errors"
)

// FetchAll is the model function which interfaces with the DB and returns a []byte of the category in json format.
func FetchAll() ([]byte, error) {

	var categories []categoryModel

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
func Create(b []byte) ([]byte, error) {

	var category categoryModel

	err := json.Unmarshal(b, &category)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	db.Save(&category)

	return []byte("Category successfully created"), nil
}

// FetchSingle gets a single todo based on param passed, returning []byte and error
func FetchSingle(id string) ([]byte, error) {

	var category categoryModel
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
func Update(b []byte, id string) ([]byte, error) {

	var category, updatedCategory categoryModel
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
	db.Model(&category).Update("category_image_url", updatedCategory.CategoryImageURL)
	db.Model(&category).Update("category_description", updatedCategory.CategoryDescription)
	db.Model(&category).Update("feeds_count", updatedCategory.FeedsCount)		

	js, err := json.Marshal(&category)
	if err != nil {
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// Delete deletes the categoryo from the database
func Delete(id string) ([]byte, error) {

	var category categoryModel
	db.First(&category, id)

	if category.ID == 0 {
		// w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte("Todo not found"))
		// return
	}

	db.Delete(&category)

	js, err := json.Marshal(&category)
	if err != nil {
		panic("Unable to marshal category into json")
	}

	return js, nil
}