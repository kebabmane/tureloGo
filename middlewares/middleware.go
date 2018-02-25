package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mendsley/gojwk"
)

// CheckJWT does the auth0 dance
func CheckJWT() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		myToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
		log.Println("the token:", myToken)
		token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return myLookupKey(token.Header["kid"].(string))
		})
		if err != nil {
			log.Println(err)
			log.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			_, err = w.Write([]byte("Unauthorized"))
			if err != nil {
				log.Printf("%s", err)
			}
		} else {
			// if the err isnt caught then token must be good so let them pass
			log.Println("Token is valid, you shall pass!")
			next.ServeHTTP(w, r)
		}
	}
}

func myLookupKey(kid string) (interface{}, error) {
	log.Printf("Kid : %v\n", kid)
	var v map[string]interface{}
	parseJSONFromURL("https://apac-syd-partner01-test.apigee.net/.well-known/openid-configuration", &v)
	var keys struct{ Keys []gojwk.Key }
	parseJSONFromURL(v["jwks_uri"].(string), &keys)
	for _, key := range keys.Keys {
		if key.Kid == kid {
			return key.DecodePublicKey()
		}
	}
	return nil, fmt.Errorf("Key not found")
}

func parseJSONFromURL(url string, v interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%s", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("%s", err)
		}
		json.Unmarshal(body, v)
	}
}

func checkUserID(userID int) bool {
	// Check your database to make sure the provided ID is valid.
	return true // or false
}
