package errors

import "errors"

type AuthError interface {
	Error() string
}

var (
	ErrUserNotExist = errors.New("user with provided username doesn't exist")
	ErrIncorrectPassword = errors.New("password provided is wrong")
	ErrPasswordTooShort = errors.New("password provided is too short. Minimal length is 8")

	ErrUsernameAlreadyRegistered = errors.New("user with such a username is already registered")
)
