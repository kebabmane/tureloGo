package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/auth0-community/auth0"
	"github.com/codegangsta/negroni"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var emailToProfileIDCache map[string]int64

// CheckJWT does the auth0 dance
func CheckJWT() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		jwksURI := "https://" + os.Getenv("AUTH0_DOMAIN") + "/.well-known/jwks.json"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwksURI})
		aud := "https://" + os.Getenv("AUTH0_AUDIENCE")
		audience := []string{aud}

		auth0ApiIssuer := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
		configuration := auth0.NewConfiguration(client, audience, auth0ApiIssuer, jose.RS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

// CheckScope is required to peek into the JWT and see wtf is going on
func CheckScope(r *http.Request, validator *auth0.JWTValidator, token *jwt.JSONWebToken) bool {
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
