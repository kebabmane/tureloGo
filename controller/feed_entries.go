package controller

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/model"
)

// FetchAllFeedEntries fetches from model and returns json
func FetchAllFeedEntries(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.FetchAllFeedEntries(id)

	handleErrorAndRespond(js, err, w)
}

// CreateFeedEntry takes request body and sends it to model, sending back success message or error on response
func CreateFeedEntry(w http.ResponseWriter, r *http.Request) {
	// read stuff from the request
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.CreateFeedEntry(b)

	handleErrorAndRespond(js, err, w)
}

// FetchSingleFeedEntry takes URL param and passes to model,
func FetchSingleFeedEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.FetchSingleFeedEntry(id)

	handleErrorAndRespond(js, err, w)
}

// UpdateFeedEntry modifies the content of Todo based on url param and body content.
func UpdateFeedEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.UpdateFeed(b, id)

	handleErrorAndRespond(js, err, w)
}

// DeleteFeedEntry deletes a feed
func DeleteFeedEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.DeleteFeedEntry(id)

	handleErrorAndRespond(js, err, w)
}
