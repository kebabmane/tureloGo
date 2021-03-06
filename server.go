package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/caarlos0/env"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kebabmane/tureloGo/app"
	"github.com/kebabmane/tureloGo/config"
	"github.com/kebabmane/tureloGo/router"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// Config is a global struct
type Config struct {
	Port              int    `env:"SERVER_PORT" envDefault:"8080"`
	HealthName        string `env:"HEALTH_NAME"`
	NegroniLoggerName string `env:"NEGRONI_LOGGER_NAME"`
	VerboseLogging    bool   `env:"VERBOSE_LOGGING,required"`
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Version: %s Commit: %s Branch: %s Date: %s\n", app.Version, app.Commit, app.Branch, app.BuildDate)
		flag.PrintDefaults()
	}

	// load application configurations, ensure some vars are checked
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalln("%+v\n", err)
	}

	// log all the config vars that have been discovered
	log.Printf("%+v\n", cfg)

	if cfg.VerboseLogging {
		log.SetLevel(log.DebugLevel)
	}

	// set the log format to JSON
	log.SetFormatter(&log.JSONFormatter{})

	// migrate and setup the database object
	db := config.Init()

	defer db.Close()

	// set up router
	r := mux.NewRouter()

	// register routes under v1 api
	r.HandleFunc("/", HomeHandler)

	// setup a subrouter
	v1 := mux.NewRouter()

	// setup the path prefix and put it through the CheckJWT middleware function
	r.PathPrefix("/v1").Handler(negroni.New(
		negroni.NewRecovery(),
		//negroni.HandlerFunc(middlewares.CheckJWT()),
		negroni.NewLogger(),
		negroni.Wrap(v1),
	))

	// setup the subrouter and then head on over to register APIs
	v1apis := v1.PathPrefix("/v1").Subrouter()
	router.RegisterHandlers(v1apis)

	// setup default CORS config
	handler := cors.Default().Handler(r)

	// start the server
	address := fmt.Sprintf(":%v", cfg.Port)
	log.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, handlers.RecoveryHandler()(handler)))

}

// HomeHandler is a sample function to show how you would use the HandleFunc
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{\"message\": \"HELLO WORLD!!\"}"))
	if err != nil {
		log.Println("cannot hello wold :(")
	}
}
