package controller

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/model"
)

// FetchAllCategories fetches from model and returns json
func FetchAllCategories(w http.ResponseWriter, r *http.Request) {

	js, err := model.FetchAllCategories()

	handleErrorAndRespond(js, err, w)
}

// CreateCategory takes request body and sends it to model, sending back success message or error on response
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// read stuff from the request
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.CreateCategory(b)

	handleErrorAndRespond(js, err, w)
}

// FetchSingleCategory takes URL param and passes to model,
func FetchSingleCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.FetchSingleCategory(id)

	handleErrorAndRespond(js, err, w)
}

// UpdateCategory modifies the content of Todo based on url param and body content.
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.UpdateCategory(b, id)

	handleErrorAndRespond(js, err, w)
}

// DeleteCategory deletes a category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.DeleteCategory(id)

	handleErrorAndRespond(js, err, w)
}
