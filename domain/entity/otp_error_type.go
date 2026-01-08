package entity

import (
	"errors"
)

var (
	ErrInvalidUserID = errors.New("user_id cannot be empty string")
)