package userservice

import (
	"SayLoveWall/app/models"
	"SayLoveWall/config/database"
)

func CheckUserExistByPhoneNum(phoneNum string) error {
	result := database.DB.Where("phone_num = ?", phoneNum).First(&models.Userdata{})

	return result.Error
}

func GetUserByPhoneNum(phoneNum string) (*models.Userdata, error) {
	var userdata models.Userdata
	result := database.DB.Where("phone_num = ?", phoneNum).First(&userdata)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userdata, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func Register(userdata models.Userdata) error {
	result := database.DB.Create(&userdata)
	return result.Error
}