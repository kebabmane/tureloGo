package healthz

import (
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"

	health "github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/checkers"
	"github.com/InVisionApp/go-health/loggers"
	"github.com/kebabmane/tureloGo/app"
)

// Health should be global, accessing through server.go to intiate the health checks
var Health *health.Health

type customCheck struct{}

// SetupHealthChecks Create a new health instance
func SetupHealthChecks() {
	// Create a new health instance
	Health = health.New()

	// Set the logger through the logrus instance
	Health.Logger = loggers.NewLogrus(nil)

	goodTestURL, _ := url.Parse("https://google.com")
	badTestURL, _ := url.Parse("google.com")

	// Instantiate your custom check
	cc := &customCheck{}

	// Create a couple of checks
	goodHTTPCheck, _ := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: goodTestURL,
	})

	badHTTPCheck, _ := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: badTestURL,
	})

	// Add the checks to the health instance
	Health.AddChecks([]*health.Config{
		{
			Name:     "good-check",
			Checker:  goodHTTPCheck,
			Interval: time.Duration(2) * time.Second,
			Fatal:    true,
		},
		{
			Name:     "bad-check",
			Checker:  badHTTPCheck,
			Interval: time.Duration(2) * time.Second,
			Fatal:    false,
		},
		{
			Name:     "app-details",
			Checker:  cc,
			Interval: time.Duration(2) * time.Second,
			Fatal:    true,
		},
	})

	//  Start the healthcheck process
	if err := Health.Start(); err != nil {
		log.Fatalf("Unable to start healthcheck: %v", err)
	}
}

// Satisfy the go-health.ICheckable interface
func (c *customCheck) Status() (interface{}, error) {
	// You can return additional information pertaining to the check as long
	// as it can be JSON marshalled
	return map[string]string{"App Version": app.Version, "App Commit": app.Commit, "App Branch": app.Branch, "App Build Date": app.BuildDate}, nil
}
