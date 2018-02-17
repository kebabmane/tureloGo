package controller

import (
	"net/http"
)

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
		w.Write(js)
	}

	// Handle the success case
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
