package exception

import (
	"errors"
)

var (
	Err400UsernameAlvailable      = errors.New("USERNAME ALVAILABLE")
	Err400PhoneAlvailable         = errors.New("PHONE NUMBER ALVAILABLE")
	Err400ProfileConfigAlvailable = errors.New("PROFILE CONFIG ALVAILABLE")
	Err400ProfileAlvailable       = errors.New("PROFILE ALVAILABLE")
	Err400InvalidTimeLayout       = errors.New("INVALID TIME LAYOUT 15:04")
	Err400InvalidIanaTimezone     = errors.New("INVALID IANA TIMEZONE")
	Err500TxNil                   = errors.New("TX PROPERTY IS NIL")
	Err500InvalidFormatEmail      = errors.New("INVALID EMAIL USER FROM UPDATE ACCOUNT")
)
