package store

import "errors"

var (
	ErrBadData  = errors.New("bad data")
	ErrNotFound = errors.New("not found")
)
