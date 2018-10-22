package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	user     = "root"
	password = ""
	dbname   = "gomeet"
)

var db *gorm.DB

// GetMysqlInfo returns the infomation needed to set up the database
func getMysqlInfo() string {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, dbname)
	return mysqlInfo
}

// ConnectToDatabase connects the application to the database to be
// used and then returns the databse
func ConnectToDatabase() *gorm.DB {
	mysqlInfo := getMysqlInfo()
	database, err := gorm.Open("mysql", mysqlInfo)
	fmt.Println(mysqlInfo)
	if err != nil {
		log.Fatal("Unable to open db connection: ", err)
	}
	database.LogMode(true)
	setDatabase(database)
	return database
}

// SetDatabase is responsible for setting up the databse
func setDatabase(database *gorm.DB) {
	db = database
}
