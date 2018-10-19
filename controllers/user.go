package controllers

import (
	"GoMeet/view"
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
	uc.HomeView.Render(w, nil)
}

// Login is responsible for showing the resister view
//
// GET /login
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	uc.LoginView.Render(w, nil)
}

// Register is responsible for showing the resister view
//
// GET /register
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	uc.RegisterView.Render(w, nil)
}
