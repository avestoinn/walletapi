package errors

import "errors"

var (

	ErrNoUsernameProvided = errors.New("no username provided in the form")  // UserCreating
	ErrNoPasswordProvided = errors.New("no password provided in the form")


	ErrNoBalanceProvided = errors.New("no balance value provided in the form")	// Balance adding and setting
)