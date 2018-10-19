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
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Unable to serve: ", err)
	}
}
