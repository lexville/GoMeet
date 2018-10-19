package models

import (
	"GoMeet/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func Connect() *UserModel {
	mysqlInfo := config.GetMysqlInfo()
	db, err := gorm.Open("mysql", mysqlInfo)
	fmt.Println(mysqlInfo)
	if err != nil {
		log.Fatal("Unable to open db connection: ", err)
	}
	db.LogMode(true)
	return &UserModel{
		db: db,
	}
}

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null;unique_index"`
	Hash     string `gorm:"not null"`
}

func (um *UserModel) Create(user *User) error {
	user.Hash = ""
	return um.db.Create(user).Error
}

func (um *UserModel) AutoMigrate() error {
	return um.db.AutoMigrate(&User{}).Error
}

func (um *UserModel) DropTable() *gorm.DB {
	return um.db.DropTableIfExists(&User{})
}
