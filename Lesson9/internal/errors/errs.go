package errors

import "errors"

var (
	ErrValidationFailed     = errors.New("validation failed")
	ErrSomethingWentWrong   = errors.New("something went wrong")
	ErrAccountNotFound      = errors.New("account not found")
	ErrUserNotFound         = errors.New("user not found")
	ErrNotEnoughBalance     = errors.New("not enough balance")
	ErrInvalidOperationType = errors.New("invalid operation type")
)
