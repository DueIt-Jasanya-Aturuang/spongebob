package usecase

import (
	"errors"
)

var ProfileNotFound = errors.New("profile tidak ditemukan")
var ProfileUserIDAndReqUserIDNotMatch = errors.New("profile anda terlarang")
var InvalidTime = errors.New("invaid time")
var ProfileConfigIsExist = errors.New("anda sudah memiliki profile config")
var ProfileConfigNotFound = errors.New("profile config tidak ditemukan")
var UserNotFound = errors.New("user tidak ditemukan")
var PhoneNumberIsExist = errors.New("nomer hp sudah terdaftar")
var NotificationNotFound = errors.New("notifikasi tidak ditemukan")
