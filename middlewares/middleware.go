package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/auth0-community/auth0"
	"github.com/codegangsta/negroni"
	raven "github.com/getsentry/raven-go"
	"github.com/kebabmane/tureloGo/config"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var emailToProfileIDCache map[string]int64

// CheckJWT does the auth0 dance
func CheckJWT() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// get environment config
		config := config.GetConfig()

		jwksURI := "https://" + config.GetString("auth0.domain") + "/.well-known/jwks.json"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwksURI})
		aud := config.GetString("auth0.audience")
		audience := []string{aud}

		auth0ApiIssuer := "https://" + config.GetString("auth0.domain") + "/"
		configuration := auth0.NewConfiguration(client, audience, auth0ApiIssuer, jose.RS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			raven.CaptureErrorAndWait(err, nil)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			fmt.Println("Token is valid, you shall pass!")
			next.ServeHTTP(w, r)
		}
	}
}

func checkScope(r *http.Request, validator *auth0.JWTValidator, token *jwt.JSONWebToken) bool {
	claims := map[string]interface{}{}
	err := validator.Claims(r, token, &claims)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if claims["scope"] != nil && strings.Contains(claims["scope"].(string), "read:messages") {
		return true
	} else {
		return false
	}
}
