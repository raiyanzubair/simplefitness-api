package model

import "errors"

// Common errors
var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrItemNotFound        = errors.New("Item not found")
	ErrItemConflict        = errors.New("Item already exists")
	Err                    = errors.New("Invalid type")
)
