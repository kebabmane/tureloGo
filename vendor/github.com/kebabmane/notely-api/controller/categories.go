package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/model"
)

// FetchAllCategories fetches from model and returns json
func FetchAllCategories(w http.ResponseWriter, r *http.Request) {

	js, err := model.FetchAllCategories()

	w.Header().Set("Content-Type", "application/json")

	fmt.Println("this is the error:", err)
	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write(js)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// CreateCategory takes request body and sends it to model, sending back success message or error on response
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// read stuff from the request
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.CreateCategory(b)

	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write(js)
		} else if err.Error() == "Bad request" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please check your inputs and try again"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sorry, something went wrong."))
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category successfully created!"))
}

// FetchSingleCategory takes URL param and passes to model,
func FetchSingleCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.FetchSingleCategory(id)

	if err != nil {
		panic("Unable to convert category to JSON format")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// UpdateCategory modifies the content of Todo based on url param and body content.
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.UpdateCategory(b, id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Todo not found"))
		} else if err.Error() == "Malformed input" {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("Please check your inputs and try again."))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong."))
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// DeleteCategory deletes a category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.DeleteCategory(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Category not found"))
		} else if err.Error() == "Unable to marshal category into json" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong."))
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
