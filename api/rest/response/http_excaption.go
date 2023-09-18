package response

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"net/http"
)

// Err500 ErrHTP internal server error
func Err500(msg string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Err:     err,
	}
}

// Err422 ErrHTP internal server error
func Err422(msg map[string]string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
		Err:     err,
	}
}

// Err409 ErrHTP internal server error
func Err409(msg map[string][]string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusConflict,
		Message: msg,
		Err:     err,
	}
}

// Err408 ErrHTP internal server error
func Err408(msg string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusRequestTimeout,
		Message: msg,
		Err:     err,
	}
}

// Err404 ErrHTP internal server error
func Err404(msg string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusNotFound,
		Message: msg,
		Err:     err,
	}
}

// Err403 ErrHTP internal server error
func Err403(msg string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusForbidden,
		Message: msg,
		Err:     err,
	}
}

// Err401 ErrHTP internal server error
func Err401(msg string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusUnauthorized,
		Message: msg,
		Err:     err,
	}
}

// Err400 ErrHTP internal server error
func Err400(msg map[string][]string, err error) error {
	return &model.ErrHTTP{
		Code:    http.StatusBadRequest,
		Message: msg,
		Err:     err,
	}
}
