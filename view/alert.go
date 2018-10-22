package view

import (
	"log"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

const (
	// AlertError is an alert label also serves as a class name
	AlertError = "alert-danger"
	// AlertWarning is an alert label also serves as a class name
	AlertWarning = "alert-warning"
	// AlertInfo is an alert label also serves as a class name
	AlertInfo = "alert-info"
	// AlertSuccess is an alert label also serves as a class name
	AlertSuccess = "alert-success"
)

// Alert contains the level of the alert
// as well as the alert message
type Alert struct {
	AlertLevel   string
	AlertMessage string
}

// RedirectWithAlert is responsible for redirecting once a specific action is done and persisting
// alerts so that they can be used in the new page
func RedirectWithAlert(w http.ResponseWriter, r *http.Request, url string, code int, alert Alert) {
	persistAlert(w, alert)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// UserCookie contains both the session as well
// as the username
type UserCookie struct {
	Session  string
	UserName string
}

// RedirectWithUserSession is responsible for redirecting once a specific action is done and persisting
// alerts so that they can be used in the new page
func RedirectWithUserSession(w http.ResponseWriter, r *http.Request, url string, code int, name string) {
	sessionID, err := uuid.NewV4()
	if err != nil {
		log.Fatal("Unable to generate a session id: ", err)
	}
	sessionCookie := http.Cookie{
		Name:  "user_session",
		Value: sessionID.String(),
	}
	userNameCookie := http.Cookie{
		Name:  "username",
		Value: name,
	}
	http.SetCookie(w, &sessionCookie)
	http.SetCookie(w, &userNameCookie)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func persistAlert(w http.ResponseWriter, alert Alert) {
	expiresAt := time.Now().Add(5 * time.Second)
	alertLevel := http.Cookie{
		Name:     "alert_level",
		Value:    alert.AlertLevel,
		Expires:  expiresAt,
		HttpOnly: true,
	}
	alertMessage := http.Cookie{
		Name:     "alert_message",
		Value:    alert.AlertMessage,
		Expires:  expiresAt,
		HttpOnly: true,
	}
	http.SetCookie(w, &alertLevel)
	http.SetCookie(w, &alertMessage)
}

func getAlert(r *http.Request) *Alert {
	alertLevel, err := r.Cookie("alert_level")
	if err != nil {
		return nil
	}
	alertMessage, err := r.Cookie("alert_message")
	if err != nil {
		return nil
	}
	alert := Alert{
		AlertLevel:   alertLevel.Value,
		AlertMessage: alertMessage.Value,
	}
	return &alert
}

func getSession(r *http.Request) *UserCookie {
	session, err := r.Cookie("user_session")
	if err != nil {
		return nil
	}
	username, err := r.Cookie("username")
	if err != nil {
		return nil
	}
	userSession := UserCookie{
		Session:  session.Value,
		UserName: username.Value,
	}
	return &userSession
}

func clearAlert(w http.ResponseWriter) {
	alertLevel := http.Cookie{
		Name:     "alert_level",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	alertMessage := http.Cookie{
		Name:     "alert_message",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(w, &alertLevel)
	http.SetCookie(w, &alertMessage)
}
