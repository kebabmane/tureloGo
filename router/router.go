package router

import (
	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/controller"
)

var err error

// RegisterHandlers and it's sub router
func RegisterHandlers(r *mux.Router) {
	r.StrictSlash(true)

	// feed routes
	r.HandleFunc("/feeds/", controller.FetchAllFeeds).Methods("GET")
	r.HandleFunc("/feeds/", controller.CreateFeed).Methods("POST")
	r.HandleFunc("/feeds/{id}", controller.FetchSingleFeed).Methods("GET")
	r.HandleFunc("/feeds/{id}", controller.UpdateFeed).Methods("PUT")
	//r.HandleFunc("/feeds/{id}/feedEntries", controller.FetchAllFeedEntries).Methods("GET")
}
