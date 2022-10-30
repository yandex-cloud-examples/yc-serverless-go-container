package main

import (
	"fmt"
	"net/http"
)

var _ error = &UserError{}

func NotFound(path string) error {
	return &UserError{
		httpCode: 404,
		message:  fmt.Sprintf("resource not found: %s", path),
	}
}
func BadRequest(msg string) error {
	return &UserError{
		httpCode: http.StatusBadRequest,
		message:  fmt.Sprintf("bad request: %s", msg),
	}
}

type UserError struct {
	httpCode int
	message  string
}

func (u *UserError) GetHTTPCode() int {
	return u.httpCode
}

func (u *UserError) Error() string {
	return u.message
}
