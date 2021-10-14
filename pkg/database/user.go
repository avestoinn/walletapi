package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"walletapi/pkg/errors"
)

type User struct {
	gorm.Model	`json:"-"`
	Username string	`json:"username"`
	Password string	`json:"-"`
	Wallets []*Wallet	`gorm:"many2many:user_wallets;" json:"wallets"`
	Token string	`json:"-"`
}


func (u *User) GetPrimaryWallet() *Wallet {
	wallet, err := GetWalletByID(u.Wallets[0].ID)
	if err != nil {
		return nil
	}
	return wallet
}


func (u *User) AddBalance(amount float64) (oldBalance, newBalance float64) {
	oldBalance = u.GetPrimaryWallet().Balance
	u.GetPrimaryWallet().Balance += amount
	database.Save(u)
	return oldBalance, u.GetPrimaryWallet().Balance
}


func (u *User) SetBalance(amount float64) (oldBalance, newBalance float64) {
	oldBalance = u.GetPrimaryWallet().Balance
	u.GetPrimaryWallet().Balance = amount
	database.Save(u)
	return oldBalance, u.GetPrimaryWallet().Balance
}


func GetUserByUsername(username string) (*User, error) {
	// Trying to find a user with such a username
	var user *User
	res := database.First(&user, "username = ?", username)
	if res.Error != nil {
		return nil, errors.ErrAccountNotExist
	}
	database.Preload(clause.Associations).Find(&user)
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

	// Creating user's first wallet
	_, err := NewWallet("", []*User{user})
	if err != nil {
		return nil, err
	}

	// Setting created wallet as primary wallet for further operations
	database.Save(user)
	database.Preload(clause.Associations).Find(&user)
	return user, nil
}