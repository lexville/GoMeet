package view

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// UserCookie contains a remember token
type UserCookie struct {
	RememberToken string
}

// RedirectWithUserSession is responsible for redirecting once a specific action is done and persisting
// remember_token so that they can be used in the new page
func RedirectWithUserSession(w http.ResponseWriter, r *http.Request, url string, code int) {
	sessionID, err := uuid.NewV4()
	if err != nil {
		log.Fatal("Unable to generate a session id: ", err)
	}
	sessionCookie := http.Cookie{
		Name:  "remember_token",
		Value: sessionID.String(),
	}
	http.SetCookie(w, &sessionCookie)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func getSession(r *http.Request) *UserCookie {
	session, err := r.Cookie("remember_token")
	if err != nil {
		return nil
	}
	userSession := UserCookie{
		RememberToken: session.Value,
	}
	return &userSession
}
