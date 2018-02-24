package model

import (
	"encoding/json"
	"errors"
)

// FetchAllLabels is the model function which interfaces with the DB and returns a []byte of the category in json format.
func FetchAllLabels() ([]byte, error) {

	var labels []Label

	db.Find(&labels)

	if len(labels) <= 0 {
		err := errors.New("Not found")
		{
			return []byte("labels not found"), err
		}
	}

	js, err := json.Marshal(labels)
	{
		return js, err
	}
}

// CreateLabel creates a new category item and returns the []byte json object and an error.
func CreateLabel(b []byte) ([]byte, error) {

	var label Label

	err := json.Unmarshal(b, &label)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	db.Save(&label)

	return []byte("Label successfully created"), nil
}

// FetchSingleLabel gets a single todo based on param passed, returning []byte and error
func FetchSingleLabel(id string) ([]byte, error) {

	var label Label
	db.First(&label, id)

	if label.ID == 0 {
		err := errors.New("Not found")
		return []byte("label not found"), err
	}

	js, err := json.Marshal(label)
	if err != nil {
		js = []byte("Unable to convert label to JSON format")
	}

	return js, err
}

// UpdateLabel is the model function for PUT
func UpdateLabel(b []byte, id string) ([]byte, error) {

	var label, updatedLabel Label
	db.First(&label, id)

	if label.ID == 0 {
		err := errors.New("Not found")
		return []byte("label not found"), err
	}

	err := json.Unmarshal(b, &updatedLabel)
	if err != nil {
		return []byte("Malformed input"), err
	}

	db.Model(&label).Update("label_name", updatedLabel.LabelName)

	js, err := json.Marshal(&label)
	if err != nil {
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// DeleteLabel deletes the label from the database
func DeleteLabel(id string) ([]byte, error) {

	var label Label
	db.First(&label, id)

	if label.ID == 0 {
		err := errors.New("Not found")
		return []byte("label not found"), err
	}

	db.Delete(&label)

	js, err := json.Marshal(&label)
	if err != nil {
		panic("Unable to marshal label into json")
	}

	return js, nil
}
