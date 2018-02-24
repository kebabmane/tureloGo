package router

import (
	"github.com/InVisionApp/go-health/handlers"
	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/controller"
	"github.com/kebabmane/tureloGo/healthz"
)

var err error

// RegisterHandlers and it's sub router
func RegisterHandlers(r *mux.Router) {
	r.StrictSlash(true)

	// health check handler
	healthHandler := handlers.NewJSONHandlerFunc(healthz.Health, map[string]interface{}{})
	r.Handle("/healthz", healthHandler)

	// categories routes
	r.HandleFunc("/categories/", controller.FetchAllCategories).Methods("GET")
	r.HandleFunc("/categories/", controller.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", controller.FetchSingleCategory).Methods("GET")
	r.HandleFunc("/categories/{id}", controller.UpdateCategory).Methods("PUT")

	// feed routes
	r.HandleFunc("/feeds/", controller.FetchAllFeeds).Methods("GET")
	r.HandleFunc("/feeds/", controller.CreateFeed).Methods("POST")
	r.HandleFunc("/feeds/{id}", controller.FetchSingleFeed).Methods("GET")
	r.HandleFunc("/feeds/{id}", controller.UpdateFeed).Methods("PUT")
	r.HandleFunc("/feeds/{id}/feedEntries", controller.FetchAllFeedEntries).Methods("GET")
}
