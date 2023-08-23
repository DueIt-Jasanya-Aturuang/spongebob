package exception

import (
	"fmt"
	"net/http"
)

type ErrResp struct {
	Code    int
	Message any
	Err     error
}

func (err *ErrResp) Error() string {
	return fmt.Sprintf("%d | %v | %v", err.Code, err.Message, err.Err)
}

func Err500(msg string, err error) error {
	return &ErrResp{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Err:     err,
	}
}

func Err422(msg map[string][]string, err error) error {
	return &ErrResp{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
		Err:     err,
	}
}

func Err408(msg string, err error) error {
	return &ErrResp{
		Code:    http.StatusRequestTimeout,
		Message: msg,
		Err:     err,
	}
}

func Err404(msg string) error {
	return &ErrResp{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func Err403(msg string) error {
	return &ErrResp{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}

func Err401(msg string) error {
	return &ErrResp{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func Err400(msg map[string][]string) error {
	return &ErrResp{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}
