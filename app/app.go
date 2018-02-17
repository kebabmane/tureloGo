package app

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

// Store allows for a global store object
var (
	Store *sessions.CookieStore
)

// Init setup a secure session store in the future
func Init() error {
	Store = sessions.NewCookieStore([]byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}
