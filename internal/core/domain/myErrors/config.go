package myerrors

import "errors"

var (
	ErrConfInvalid  = errors.New("invalid conf")
	ErrConfNotFound = errors.New("env not found")
)
