package user

import (
	"cookie-shop-api/models"
	"cookie-shop-api/models/dto"
)

// 登录是否成功，登录的角色{admin|user}
func Login(req dto.LoginRequest) (bool, string, *dto.CurrentUserData, error){
	data, err := models.GetUserByUsername(req.Username)

	if data == nil {
		return false, "", nil, models.ErrNoRegister
	}

	if err != nil {
		return false,"", nil, err
	}

	if data.Password != req.Password {
		return false, "", nil, models.ErrWrongPassword
	}

	if data.Isvalidate != "1" {
		return false, "",nil, models.ErrUserInvalid
	}



	if data.Isadmin == "1" {
		return true, "admin", data.ToDtoUser() , nil
	}

	return true, "user", data.ToDtoUser(), nil
}
