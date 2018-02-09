package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MindscapeHQ/raygun4go"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/controller"
	"github.com/kebabmane/tureloGo/middlewares"
	"github.com/kebabmane/tureloGo/model"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"github.com/zbindenren/negroni-prometheus"
)

func main() {
	// load application configurations in not production

	if os.Getenv("ENV") == "PRODUCTION" {
		fmt.Println("your running in production, did you know that?")
	} else {
		fmt.Println("your running in dev/test, did you know that?")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if os.Getenv("RAYGUN_APIKEY") == "" {
		log.Printf("No Raygun API KEY found, no app tracing")
	} else {
		raygun, err := raygun4go.New("tureloGo", os.Getenv("RAYGUN_APIKEY"))
		if err != nil {
			log.Println("Unable to create Raygun client:", err.Error())
		}
		log.Printf("Pew pew - raygun tracing is enabled")
		defer raygun.HandleError()
	}

	// migrate the database
	model.Init()

	// create the logger
	logger := logrus.New()

	// CORS middleware setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type", "Origin", "Accept-Encoding", "Accept-Language", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowCredentials: true,
	})

	// set up router
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/health", controller.HealthFunction).Methods("GET")

	r.Handle("/metrics", prometheus.Handler())

	// s is a subrouter to handle question routes
	api := r.PathPrefix("/v1").Subrouter()

	// categories routes
	api.HandleFunc("/categories/", controller.FetchAllCategories).Methods("GET")
	api.HandleFunc("/categories/", controller.CreateCategory).Methods("POST")
	api.HandleFunc("/categories/{id}", controller.FetchSingleCategory).Methods("GET")
	api.HandleFunc("/categories/{id}", controller.UpdateCategory).Methods("PUT")

	// feed routes
	api.HandleFunc("/feeds", controller.FetchAllFeeds).Methods("GET")
	api.HandleFunc("/feeds/", controller.CreateFeed).Methods("POST")
	api.HandleFunc("/feeds/{id}", controller.FetchSingleFeed).Methods("GET")
	api.HandleFunc("/feeds/{id}", controller.UpdateFeed).Methods("PUT")

	// feedEntry routes
	api.HandleFunc("/feedEntries/{id}", controller.FetchAllFeedEntries).Methods("GET")

	// muxRouter uses Negroni handles the middleware for authorization
	muxRouter := http.NewServeMux()
	muxRouter.Handle("/", r)
	muxRouter.Handle("/api/", negroni.New(
		negroni.HandlerFunc(middlewares.CheckJWT()),
		negroni.Wrap(api),
	))

	// Negroni handles the middleware chaining with next
	n := negroni.Classic()

	m := negroniprometheus.NewMiddleware("tureloGo")

	// Use promethus for service stuff
	n.Use(m)

	// Use CORS
	n.Use(c)

	// handle routes with the muxRouter
	n.UseHandler(muxRouter)

	// start the server
	address := fmt.Sprintf(":%v", os.Getenv("PORT"))
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, handlers.RecoveryHandler()(n)))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"Hello world\"}"))
}
