package app

import "errors"

const (
	//Default is the default error code
	Default = 0
	//Notfound is used when a user try retrieve a resource which is not available
	Notfound = 1
	//DuplicateData is used when a user try to create something with an already existed data
	DuplicateData = 2
	//InvalidData is used when user sends invalid data
	InvalidData = 3
	//Internal is used when ther is an unhandled error
	Internal = 99
)

var (
	ErrInvalidData = errors.New("invalid Data Request")
	ErrInternal    = errors.New("internal error, please contact admin")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}
