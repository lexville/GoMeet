package routes

import (
	"GoMeet/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SetUpRoutes contains all the routes used
// by the application
func SetUpRoutes() {
	userController := controllers.AddViewTemplates()
	r := mux.NewRouter()
	r.HandleFunc("/", userController.Get).Methods("GET")
	r.HandleFunc("/register", userController.Register).Methods("GET")
	r.HandleFunc("/register", userController.Create).Methods("POST")
	r.HandleFunc("/login", userController.Login).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Unable to serve: ", err)
	}
}
