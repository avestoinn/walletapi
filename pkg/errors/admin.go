package errors

import "errors"

var (
	// UserCreating
	ErrNoUsernameProvided = errors.New("no username provided in the form")
	ErrNoPasswordProvided = errors.New("no password provided in the form")

	// Balance adding and setting
	ErrNoBalanceProvided = errors.New("no balance value provided in the form")
)