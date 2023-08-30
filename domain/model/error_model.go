package model

import (
	"errors"
	"fmt"
)

// ResponseError http error response
type ResponseError struct {
	Errors     *any  `json:"errors,omitempty"`
	ErrorsReal error `json:"errors_real,omitempty"`
}

// ResponseSuccess http success response
type ResponseSuccess struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// ErrHTTP error http
type ErrHTTP struct {
	Code    int
	Message any
	Err     error
}

func (err *ErrHTTP) Error() string {
	return fmt.Sprintf("%d | %v | %v", err.Code, err.Message, err.Err)
}

var (
	ErrConflict        = errors.New("your Item already exist")
	ErrBadInput        = errors.New("your Item already exist")
	ErrUnauthorization = errors.New("UNAUTHORIZATION")
	ErrForbidden       = errors.New("FORBIDDEN")
)
