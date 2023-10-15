package userservice

import (
	"SayLoveWall/app/models"
	"SayLoveWall/config/database"
)

func CheckUserExistByPhoneNum(phoneNum string) error {
	result := database.DB1.Where("phone_num = ?", phoneNum).First(&models.Userdata{})

	return result.Error
}

func GetUserByPhoneNum(phoneNum string) (*models.Userdata, error) {
	var userdata models.Userdata
	result := database.DB1.Where("phone_num = ?", phoneNum).First(&userdata)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userdata, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func Register(userdata models.Userdata) error {
	result := database.DB1.Create(&userdata)
	return result.Error
}

func UpdateUserData(user models.Userdata) error {
	// 与课上不同，不多接收 owner_id 参数，选择使用 omit 忽略该字段的更新
	result := database.DB1.Omit("id").Save(&user)
	return result.Error
}
