package domainerror

import (
	"errors"
)

var (
	ErrUsernameAlvailable      = errors.New("USERNAME ALVAILABLE")
	ErrPhoneAlvailable         = errors.New("PHONE NUMBER ALVAILABLE")
	ErrProfileConfigAlvailable = errors.New("PROFILE CONFIG ALVAILABLE")
	ErrProfileAlvailable       = errors.New("PROFILE ALVAILABLE")
)
