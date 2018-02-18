package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/engine/memstore"
	"github.com/alexedwards/scs/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	raven "github.com/getsentry/raven-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kdar/logrus-cloudwatchlogs"
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/controller"
	"github.com/kebabmane/tureloGo/middlewares"
	"github.com/kebabmane/tureloGo/model"
	_ "github.com/lib/pq"
	"github.com/meatballhat/negroni-logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	logrus "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"github.com/zbindenren/negroni-prometheus"
)

var logger *logrus.Logger

func main() {
	// load application configurations in not production
	if os.Getenv("ENV") == "PRODUCTION" {
		fmt.Println("your running in production, did you know that?")
		fmt.Println("setting up cloudwatch logging")
		key := os.Getenv("AWS_ACCESS_KEY")
		secret := os.Getenv("AWS_SECRET_KEY")
		group := os.Getenv("AWS_CLOUDWATCHLOGS_GROUP_NAME")
		stream := os.Getenv("AWS_CLOUDWATCHLOGS_STREAM_NAME")

		// logs.us-east-1.amazonaws.com
		cred := credentials.NewStaticCredentials(key, secret, "")
		cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(cred)

		hook, err := logrus_cloudwatchlogs.NewHook(group, stream, cfg)
		if err != nil {
			log.Fatal(err)
		}
		// create the logger
		logger = logrus.New()

		logger.Hooks.Add(hook)
		logger.Out = ioutil.Discard
		logger.Formatter = logrus_cloudwatchlogs.NewProdFormatter()
	} else {
		fmt.Println("your running in dev/test, did you know that?")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		// create the logger
		logger = logrus.New()
	}

	// migrate and setup the database object
	model.Init()

	// setup raven/sentry for error logging
	raven.SetDSN(os.Getenv("SENTRY_API"))

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
	api.HandleFunc("/feeds/", controller.FetchAllFeeds).Methods("GET")
	api.HandleFunc("/feeds/", controller.CreateFeed).Methods("POST")
	api.HandleFunc("/feeds/{id}", controller.FetchSingleFeed).Methods("GET")
	api.HandleFunc("/feeds/{id}", controller.UpdateFeed).Methods("PUT")
	api.HandleFunc("/feeds/{id}/feedEntries", controller.FetchAllFeedEntries).Methods("GET")

	// muxRouter uses Negroni handles the middleware for authorization
	muxRouter := http.NewServeMux()
	muxRouter.Handle("/", r)
	muxRouter.Handle("/v1/api/", negroni.New(
		negroni.HandlerFunc(middlewares.CheckJWT()),
		negroni.Wrap(api),
	))

	// Negroni handles the middleware chaining with next
	n := negroni.Classic()

	m := negroniprometheus.NewMiddleware(os.Getenv("HEALTH_NAME"))

	// Use promethus for service stuff
	n.Use(m)

	// Use CORS
	n.Use(c)

	// Setup logrus middleware
	n.Use(negronilogrus.NewMiddlewareFromLogger(logger, "web"))

	// handle routes with the muxRouter
	n.UseHandler(muxRouter)

	// start the server
	address := fmt.Sprintf(":%v", os.Getenv("PORT"))
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, handlers.RecoveryHandler()(sessionManager(n))))

}
