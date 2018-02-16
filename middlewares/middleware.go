package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/auth0-community/auth0"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat/go-jwx/jwk"
	jose "gopkg.in/square/go-jose.v2"
)

var emailToProfileIDCache map[string]int64

// CheckJWT does the auth0 dance
func CheckJWT() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println("your in CheckJET()", r)

		jwksURI := os.Getenv("WELL-KNOWN-URL")
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwksURI})
		aud := os.Getenv("AUTH0_AUDIENCE")
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
			fmt.Println("Token is valid, you shall pass!")
			next.ServeHTTP(w, r)
		}
	}
}

func getCognitoJwk(token *jwt.Token) (interface{}, error) {

	// TODO: cache response so we don't have to make a request every time
	// we want to verify a JWT
	set, err := jwk.FetchHTTP(os.Getenv("WELL-KNOWN-URL"))
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	if key := set.LookupKeyID(keyID); len(key) == 1 {
		return key[0].Materialize()
	}

	return nil, errors.New("unable to find key")
}
