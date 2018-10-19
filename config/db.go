package config

import (
	"fmt"
)

const (
	user     = "root"
	password = ""
	dbname   = "gomeet"
)

func GetMysqlInfo() string {
	// mysqlInfo := "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	mysqlInfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, dbname)
	return mysqlInfo
}
