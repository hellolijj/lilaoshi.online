package visiblerror

import "errors"

var (
	ErrRegisterAlreadyExists = errors.New("当前用户已注册")
	ErUnknow = errors.New("未知错误")
)
