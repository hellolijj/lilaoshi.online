package models

import "errors"

var (
	ErrRegisterAlreadyExists = errors.New("当前用户已注册")
	ErrNoRegister = errors.New("当前用户没有注册，请先注册")
	ErrWrongPassword = errors.New("账号密码不匹配")
	ErrUserInvalid = errors.New("当前用户已失效")
	ErrUnknown = errors.New("未知错误")
)
