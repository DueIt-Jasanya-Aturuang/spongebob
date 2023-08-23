package exception

import (
	"errors"
)

var (
	Err400UsernameAlvailable      = errors.New("USERNAME ALVAILABLE")
	Err400PhoneAlvailable         = errors.New("PHONE NUMBER ALVAILABLE")
	Err400ProfileConfigAlvailable = errors.New("PROFILE CONFIG ALVAILABLE")
	Err400ProfileAlvailable       = errors.New("PROFILE ALVAILABLE")
	Err500TxNil                   = errors.New("TX PROPERTY IS NIL")
)
