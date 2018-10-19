package routes

import (
	"GoMeet/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes contains all the routes used
// by the application
func InitRoutes() {
	homeController := controllers.AddViewTemplates()
	r := mux.NewRouter()
	r.HandleFunc("/", homeController.Get).Methods("GET")
	r.HandleFunc("/register", homeController.Register).Methods("GET")
	r.HandleFunc("/login", homeController.Login).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Unable to serve: ", err)
	}
}