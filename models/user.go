package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

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

// CreateUser is responsible for creating a new
// user. It returns nil if the user is created
// and an error if there user isn't created
func CreateUser(user *User) error {
	paswordByte := []byte(user.Hash + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(
		paswordByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Hash = string(hashedBytes)
	return db.Create(user).Error
}

// AuthenticateUser checks in the db whether the values provided belong
// to a user in the db
func AuthenticateUser(email, password string) (*User, error) {
	foundUser, err := FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

// FindByEmail (TODO implement this)
func FindByEmail(email string) (*User, error) {
	var user User
	return &user, nil
}

// AutoMigrateUserTable migrates a user table
func AutoMigrateUserTable() error {
	return db.AutoMigrate(&User{}).Error
}

// DropUserTable destroys the user table
func DropUserTable() *gorm.DB {
	return db.DropTableIfExists(&User{})
}
