package state

import "errors"

var (
	ErrDataNotFound = errors.New(`data found`)
	ErrSomethingWentWrong = errors.New(`something went wrong`)
)
