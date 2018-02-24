package controller

import (
	"net/http"
)

// HealthFunction fetches from model and returns json
func HealthFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"Health is A OK!!\"}"))
}
