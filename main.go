package main

import (
	"GoMeet/models"
	"GoMeet/routes"
	"log"
)

func init() {
	intializeDatabases()
}

func intializeDatabases() {
	us := models.Connect()
	// us.DropTable()
	if err := us.AutoMigrate(); err != nil {
		log.Fatal("Unable to migrate the user table: ", err)
	}
}

func main() {
	routes.InitRoutes()
}
