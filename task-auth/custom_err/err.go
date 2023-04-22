package custom_err

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ExsistsUser     = errors.New("username already taken")
)
