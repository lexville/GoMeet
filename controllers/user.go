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
	service      *models.UserModel
}

// AddViewTemplates is responsible for getting all
// the templates that will be used by the user
// controller
func AddViewTemplates(us *models.UserModel) *UserController {
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
		service: us,
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
	if err := uc.service.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
