package models

import (
	"GoMeet/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// UserModel contains the db to be
// used by the model
type UserModel struct {
	db *gorm.DB
}

// UserService opens up a db connection
func UserService() *UserModel {
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

// User contains all the fields associated with the
// user
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null;unique_index"`
	Hash     string `gorm:"not null"`
}

const userPwPepper = "tZXMdcNWU5jLj57JOlcE"

// Create is responsible for creating a new
// user. It returns nil if the user is created
// and an error if there user isn't created
func (um *UserModel) Create(user *User) error {
	paswordByte := []byte(user.Hash + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(
		paswordByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Hash = string(hashedBytes)
	return um.db.Create(user).Error
}

// AutoMigrate migrates a user table
func (um *UserModel) AutoMigrate() error {
	return um.db.AutoMigrate(&User{}).Error
}

// DropTable destroys the user table
func (um *UserModel) DropTable() *gorm.DB {
	return um.db.DropTableIfExists(&User{})
}
