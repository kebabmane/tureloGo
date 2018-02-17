package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/alexedwards/scs/engine/memstore"
	"github.com/alexedwards/scs/session"
	"github.com/codegangsta/negroni"
	raven "github.com/getsentry/raven-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/config"
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

	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*enviroment)

	// get environment config
	config := config.GetConfig()

	// migrate the database
	model.Init()

	// create the logger
	logger := logrus.New()

	// setup raven/sentry for error logging
	raven.SetDSN(config.GetString("sentry.raven"))

	// setup session store
	engine := memstore.New(30 * time.Minute)
	sessionManager := session.Manage(engine, session.IdleTimeout(30*time.Minute), session.Persist(true), session.Secure(true))

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
	api := r.PathPrefix("/v1/api").Subrouter()

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
	muxRouter.Handle("/v1/api/", negroni.New(
		negroni.HandlerFunc(middlewares.CheckJWT()),
		negroni.Wrap(api),
	))

	// Negroni handles the middleware chaining with next
	n := negroni.Classic()

	m := negroniprometheus.NewMiddleware(config.GetString("health.name"))

	// Use promethus for service stuff
	n.Use(m)

	// Use CORS
	n.Use(c)

	// handle routes with the muxRouter
	n.UseHandler(muxRouter)

	// start the server
	address := fmt.Sprintf(":%v", config.GetString("server.port"))
	logger.Infof("server %v is started at %v\n", app.Version, config.GetString("server.port"))
	panic(http.ListenAndServe(address, handlers.RecoveryHandler()(sessionManager(n))))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"Hello world\"}"))
}
