package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/model"
)

// FetchAllFeeds fetches from model and returns json
func FetchAllFeeds(w http.ResponseWriter, r *http.Request) {

	js, err := model.FetchAllFeeds()

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

// CreateFeed takes request body and sends it to model, sending back success message or error on response
func CreateFeed(w http.ResponseWriter, r *http.Request) {
	// read stuff from the request
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.CreateFeed(b)

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
	w.Write([]byte("Feed successfully created!"))
}

// FetchSingleFeed takes URL param and passes to model,
func FetchSingleFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.FetchSingleFeed(id)

	if err != nil {
		panic("Unable to convert feed to JSON format")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// UpdateFeed modifies the content of Todo based on url param and body content.
func UpdateFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.UpdateFeed(b, id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Feed not found"))
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

// DeleteFeed deletes a feed
func DeleteFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.DeleteFeed(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Feed not found"))
		} else if err.Error() == "Unable to marshal feed into json" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong."))
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
