package adminservice

import (
	"SayLoveWall/app/models"
	"SayLoveWall/config/database"
)

func CheckUserExistByAccount(account string) error {
	result := database.DB3.Where("account = ?", account).First(&models.Admindata{})

	return result.Error
}

func GetAdminByAccount(account string) (*models.Admindata, error) {
	var admindata models.Admindata
	result := database.DB1.Where("account = ?", account).First(&admindata)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admindata, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}


func UpdateTextData(text models.Textdata) error {
	result := database.DB2.Save(&text)
	return result.Error
}