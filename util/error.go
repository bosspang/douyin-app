package util

import "errors"

const (
	Success = iota

	WrongAuth
	NoAuth

	UserExisted
	UserNotExist
	WrongPassword
	IsFollow
	NotFollow
	InternalServerError
	ParamError
)

var ErrCode2Msg = map[int]string{
	Success: "success",

	WrongAuth: "the token is expired or invalid",
	NoAuth:    "token not received",

	UserExisted:   "user already exists",
	UserNotExist:  "username is not exist",
	WrongPassword: "password is wrong",
	IsFollow:      "该关注记录已存在",
	NotFollow:     "该关注记录尚未存在",

	InternalServerError: "internal server error",
	ParamError:          "something wrong with param...",
}

var (
	ErrWrongAuth     = errors.New(ErrCode2Msg[WrongAuth])
	ErrNoAuth        = errors.New(ErrCode2Msg[NoAuth])
	ErrUserExisted   = errors.New(ErrCode2Msg[UserExisted])
	ErrWrongPassword = errors.New(ErrCode2Msg[WrongPassword])
	ErrIsFollow      = errors.New(ErrCode2Msg[IsFollow])
	ErrNotFollow     = errors.New(ErrCode2Msg[NotFollow])
)
