package utils

import (
	"errors"
)

var visibleError map[string]error

func init() {
	visibleError = map[string]error{
		"ErrRegisterAlreadyExists": errors.New("当前用户已经注册"),
	}
}


func responseError(err error)  {

	for _, v := range visibleError {
		if err == v {

			// 错误对用户可见

		}
	}


	
}



