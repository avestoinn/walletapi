package database

import (
	"gorm.io/gorm"
	"walletapi/pkg/errors"
)

type User struct {
	gorm.Model	`json:"-"`
	Username string	`json:"username"`
	Password string	`json:"-"`
	Balance	float64	`json:"balance"`
	Token string	`json:"-"`
}


func (u *User) AddBalance(amount float64) (oldBalance, newBalance float64) {
	oldBalance = u.Balance
	u.Balance += amount
	database.Save(u)
	return oldBalance, u.Balance
}


func (u *User) SetBalance(amount float64) (oldBalance, newBalance float64) {
	oldBalance = u.Balance
	u.Balance = amount
	database.Save(u)
	return oldBalance, u.Balance
}




func GetUserByUsername(username string) (*User, error) {
	// Trying to find a user with such a username
	var user *User
	res := database.First(&user, "username = ?", username)
	if res.Error != nil {
		return nil, errors.ErrAccountNotExist
	}
	return user, nil
}



// TryAuth Gets an instance of user if matches
func TryAuth(username, password string) (*User, error){
	// Trying to find a user with such a username
	user, err := GetUserByUsername(username)
	if user == nil {
		return nil, err
	}

	// Verifying passwords (provided and original)
	if user.Password != password {
		return nil, errors.ErrIncorrectPassword
	}

	return user, nil
}


// CreateUser creates a new instance of user if provided username doesn't exist
func CreateUser(username, password string) (*User, error) {
	// If provided password is short
	if len(password) < 8 {
		return nil, errors.ErrPasswordTooShort
	}

	// Check if there is registered user with such a username
	user, _ := GetUserByUsername(username)
	if user != nil {
		return nil, errors.ErrUsernameAlreadyRegistered
	}

	user = &User{Username: username, Password: password}
	res := database.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}