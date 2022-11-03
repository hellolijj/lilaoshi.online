package user

import (
	"cookie-shop-api/models"
	"cookie-shop-api/models/dto"
	"github.com/beego/beego/v2/client/orm"
)

func boolToString(admin bool) string {
	if admin {
		return "1"
	}
	return "0"
}

func Register(req dto.RegisterRequest) (bool, error) {
	data, err := models.GetUserByUsername(req.Username)
	if err != nil && err != orm.ErrNoRows {
		return false, err
	}

	if data != nil {
		return false, models.ErrRegisterAlreadyExists
	}

	_, err = models.AddUser(&models.User{
		Username:   req.Username,
		Password:   req.Password,
		Name:       req.Name,
		Email:      "",
		Phone:      req.Mobile,
		Address:    "",
		Isadmin:   boolToString(req.AsAdmin),
		Isvalidate: "1",
	})

	if err != nil {
		return false, err
	}

	return true, nil
}