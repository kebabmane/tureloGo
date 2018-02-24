package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

var err error

func handleErrorAndRespond(js []byte, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	// handle the error cases
	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == "Error saving to database" {
			w.WriteHeader(http.StatusInternalServerError)
		} else if err.Error() == "Something went wrong" {
			w.WriteHeader(http.StatusInternalServerError)
		} else if err.Error() == "Update error" {
			w.WriteHeader(http.StatusInternalServerError)
		} else if err.Error() == "User already exists" {
			w.WriteHeader(http.StatusFound)
		} else if err.Error() == "Unauthorized" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		// in any case, send back the message created in the model
		_, err = w.Write(js)
		if err != nil {
			log.Println("an error happened whilst handling an error")
		}
		return
	}

	// Handle the success case
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(js)
	if err != nil {
		log.Println("an error happened whilst handling an error")
	}
}
