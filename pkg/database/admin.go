package database

import (
	"gorm.io/gorm"
	"walletapi/pkg/errors"
)

type Admin struct {
	gorm.Model	`json:"-"`
	Username string	`json:"username"`
	Password string	`json:"-"`
	Token string	`json:"-"`
}


func getAdminByUsername(username string) (*Admin, error) {
	// Trying to find an admin with such a username
	var admin *Admin
	res := database.First(&admin, "username = ?", username)
	if res.Error != nil {
		return nil, errors.ErrAccountNotExist
	}
	return admin, nil
}


// TryAuthAdmin Gets an instance of user if matches
func TryAuthAdmin(username, password string) (*Admin, error){
	// Trying to find an admin with such a username
	admin, err := getAdminByUsername(username)
	if admin == nil {
		return nil, err
	}

	// Verifying passwords (provided and original)
	if admin.Password != password {
		return nil, errors.ErrIncorrectPassword
	}

	return admin, nil
}


// CreateAdmin creates a new instance of admin if provided username doesn't exist
func CreateAdmin(username, password string) (*Admin, error) {
	// If provided password is short
	if len(password) < 8 {
		return nil, errors.ErrPasswordTooShort
	}

	// Check if there is registered admin with such a username
	admin, _ := getAdminByUsername(username)
	if admin != nil {
		return nil, errors.ErrUsernameAlreadyRegistered
	}

	admin = &Admin{Username: username, Password: password}
	res := database.Create(admin)
	if res.Error != nil {
		return nil, res.Error
	}

	return admin, nil
}