package errors

import "errors"

var (
	ErrWalletNotExist = errors.New("wallet with provided id doesn't exist")
	ErrWalletWithoutOwner	= errors.New("wallet can't be created without owner")
)
