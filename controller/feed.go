package controller

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/model"
)

// FetchAllFeeds fetches from model and returns json
func FetchAllFeeds(w http.ResponseWriter, r *http.Request) {

	js, err := model.FetchAllFeeds()

	handleErrorAndRespond(js, err, w)
}

// CreateFeed takes request body and sends it to model, sending back success message or error on response
func CreateFeed(w http.ResponseWriter, r *http.Request) {
	// read stuff from the request
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.CreateFeed(b)

	handleErrorAndRespond(js, err, w)
}

// FetchSingleFeed takes URL param and passes to model,
func FetchSingleFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.FetchSingleFeed(id)

	handleErrorAndRespond(js, err, w)
}

// UpdateFeed modifies the content of Todo based on url param and body content.
func UpdateFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.UpdateFeed(b, id)

	handleErrorAndRespond(js, err, w)
}

// DeleteFeed deletes a feed
func DeleteFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.DeleteFeed(id)

	handleErrorAndRespond(js, err, w)
}
