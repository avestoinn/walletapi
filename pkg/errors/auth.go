package errors

import "errors"


var (
	ErrAccountNotExist = errors.New("account with provided username doesn't exist")
	ErrIncorrectPassword = errors.New("password provided is wrong")
	ErrPasswordTooShort = errors.New("password provided is too short. Minimal length is 8")

	ErrUsernameAlreadyRegistered = errors.New("account with such a username is already registered")
)
