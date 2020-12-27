package session

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/securecookie"

	"github.com/gorilla/sessions"
)

var (
	//Store is sessionStore
	Store *sessions.CookieStore
)

// User is
type User struct {
	UserID        uint   `json:"id"`
	Username      string `json:"name"`
	Authenticated bool   `json:"authenticated"`
	Token         string `json:"token"`
}

//Initialize is used to initialize sessions
func Initialize() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	Store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)
	Store.Options = &sessions.Options{
		Path:     "/",
		Domain:   "lourdemathaschool.edu.in",
		MaxAge:   60 * 30,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	gob.Register(User{})

}
