package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"walletapi/pkg/errors"
)

type Wallet struct {
	gorm.Model	`json:"-"`
	Name string
	Balance	float64
	Owners []*User	`gorm:"many2many:user_wallets;" json:"-"`
}


func GetWalletByID(walletId uint) (*Wallet, error) {
	// Trying to find a user with such a username
	var wallet *Wallet
	res := database.First(&wallet, "id = ?", walletId).Find(&wallet)
	if res.Error != nil {
		return nil, errors.ErrWalletNotExist
	}
	database.Preload(clause.Associations).Find(&wallet)
	return wallet, nil
}


func NewWallet(name string, owners []*User) (*Wallet, error) {
	if name == "" {
		name = "Wallet"
	}

	if len(owners) == 0 {
		return nil, errors.ErrWalletWithoutOwner
	}

	wallet := &Wallet{
		Name:    name,
		Balance: 0,
		Owners:  owners,
	}

	database.Save(wallet)
	database.Preload(clause.Associations).Find(&wallet)
	return wallet, nil
}