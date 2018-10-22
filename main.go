package main

import (
	"GoMeet/models"
	"GoMeet/routes"
)

func main() {
	db := models.ConnectToDatabase()
	defer db.Close()
	models.AutoMigrateUserTable()
	routes.SetUpRoutes()
}
