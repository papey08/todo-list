package model

import "errors"

var (
	ErrTaskRepo     = errors.New("something wrong with task database")
	ErrTaskNotFound = errors.New("todo task with required id was not found")
	ErrInvalidTask  = errors.New("some of the fields of task are invalid")
)
