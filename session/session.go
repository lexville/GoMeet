package session

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mY-s3cr3t_Ok!@"))

var (
	NoSessionUsername = errors.New("The session doesn't contain a username")
	NoSessionUserID   = errors.New("The session doesn't contain a userID")
	ConversionError   = errors.New("Unable to convert session value")
)

type SessionDetails struct {
	Username string
	UserID   string
}

func SetUserSession(w http.ResponseWriter, r *http.Request, username string) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["username"] = username
	session.Save(r, w)
}

func GetUserSession(w http.ResponseWriter, r *http.Request) (username string) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return session.Values["username"].(string)
}
