package controllers

import (
	"GoMeet/models"
	"GoMeet/view"
	"log"
	"net/http"
)

// UserController contains the views as well as
// any other piece of data needed by the user
// controller
type UserController struct {
	HomeView     *view.View
	LoginView    *view.View
	RegisterView *view.View
}

// AddViewTemplates is responsible for getting all
// the templates that will be used by the user
// controller
func AddViewTemplates() *UserController {
	return &UserController{
		HomeView: view.AddTemplateFiles(
			"base",
			"templates/welcome.gohtml"),
		LoginView: view.AddTemplateFiles(
			"base",
			"templates/user/login.gohtml"),
		RegisterView: view.AddTemplateFiles(
			"base",
			"templates/user/register.gohtml"),
	}
}

// Get is the handler responsible for the showing the
// homepage
//
// GET /
func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	uc.HomeView.Render(w, r, nil)
}

// Login is responsible for showing the resister view
//
// GET /login
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	uc.LoginView.Render(w, r, nil)

}

// Register is responsible for showing the resister view
//
// GET /register
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	uc.RegisterView.Render(w, r, nil)
}

// Create is responsible for creating a new application user
//
// POST /register
func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("Unable to parse the register form: ", err)
	}
	if r.FormValue("password") != r.FormValue("confirm-password") {
		alert := view.Alert{
			AlertLevel:   view.AlertError,
			AlertMessage: "Passwords do not match",
		}
		view.RedirectWithAlert(w, r, "/register", http.StatusSeeOther, alert)
		return
	}
	user := models.User{
		Name:     r.FormValue("name"),
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Hash:     r.FormValue("password"),
	}
	if err := models.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	alert := view.Alert{
		AlertLevel:   view.AlertSuccess,
		AlertMessage: "Success. You can now login",
	}
	view.RedirectWithAlert(w, r, "/login", http.StatusSeeOther, alert)
	return
}

// Authenticate is used to check whether the username and password provided
// belong to a user in the db
func (uc *UserController) Authenticate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("Unable to parse the register form: ", err)
	}
	// Confirm that the password has been provided
	if r.FormValue("password") == "" {
		alert := view.Alert{
			AlertLevel:   view.AlertError,
			AlertMessage: "No password provided",
		}
		view.RedirectWithAlert(w, r, "/login", http.StatusSeeOther, alert)
		return
	}
	// Confirm that the username has been provided
	if r.FormValue("email") == "" {
		alert := view.Alert{
			AlertLevel:   view.AlertError,
			AlertMessage: "No email provided",
		}
		view.RedirectWithAlert(w, r, "/login", http.StatusSeeOther, alert)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	_, err := models.AuthenticateUser(email, password)
	if err != nil {
		alert := view.Alert{
			AlertLevel:   view.AlertError,
			AlertMessage: "No user with that username or password exists",
		}
		view.RedirectWithAlert(w, r, "/login", http.StatusSeeOther, alert)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
