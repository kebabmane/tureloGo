package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kebabmane/notely-api/model"
)

// FetchAllTodos fetches from model and returns json
func FetchAllTodos(w http.ResponseWriter, r *http.Request) {

	js, err := model.FetchAllTodos()

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

// CreateTodo takes request body and sends it to model, sending back success message or error on response
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// read stuff from the request
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.CreateTodo(b)

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
	w.Write([]byte("Todo successfully created!"))
}

// FetchSingleTodo takes URL param and passes to model,
func FetchSingleTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.FetchSingleTodo(id)

	if err != nil {
		panic("Unable to convert todo to JSON format")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// UpdateTodo modifies the content of Todo based on url param and body content.
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := []byte(buf.String())

	js, err := model.UpdateTodo(b, id)

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

// DeleteTodo deletes a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	js, err := model.DeleteTodo(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Todo not found"))
		} else if err.Error() == "Unable to marshal todo into json" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong."))
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
