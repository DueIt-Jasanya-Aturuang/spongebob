package exception

import (
	"errors"
)

var (
	Err400UsernameAvailable      = errors.New("USERNAME AVAILABLE")
	Err400PhoneAvailable         = errors.New("PHONE NUMBER AVAILABLE")
	Err400ProfileConfigAvailable = errors.New("PROFILE CONFIG AVAILABLE")
	Err400ProfileAvailable       = errors.New("PROFILE AVAILABLE")
	Err400InvalidTimeLayout      = errors.New("INVALID TIME LAYOUT 15:04")
	Err400InvalidIanaTimezone    = errors.New("INVALID IANA TIMEZONE")
	Err500TxNil                  = errors.New("TX PROPERTY IS NIL")
	Err500InvalidFormatEmail     = errors.New("INVALID EMAIL USER FROM UPDATE ACCOUNT")
	Err401Msg                    = errors.New("UNAUTHORIZATION")
)

var (
	Required     = "%s is required"
	MaxString    = "maximum %s character must be %d"
	MinString    = "minimum %s character must be %d"
	Gender       = "%s gender must be male, female, or undefinied"
	FileSize     = "max %s size should be %d kb or %d mb"
	FileContent  = "%s must be %s"
	InvalidField = "%s invalid %s, example %s"
	InvalidID    = "invalid %s"
	Integer      = "%s must be number"
	Enum         = "%s must be %s"
	MinInteger   = "minimum %s number must be %d"
	MaxInteger   = "maximum %s number must be %d"
)
