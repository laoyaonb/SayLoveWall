package textservice

import (
	"SayLoveWall/app/models"
	"SayLoveWall/config/database"
)

func WordCountJudge(str string) bool {
	return len(str) <= 900
}

func Publish(textdata models.Textdata) error {
	result := database.DB2.Create(&textdata)
	return result.Error
}

func Delete(id uint) error {
	result := database.DB2.Where("id = ?", id).Delete(&models.Textdata{})
	return result.Error
}

func GetTextPersonalList(Name string) ([]models.Textdata, error) {
	result := database.DB2.Where("name = ?", Name).First(&models.Textdata{})
	if result.Error != nil {
		return nil, result.Error
	}
	var textList []models.Textdata
	result = database.DB2.Where("name = ?", Name).Find(&textList)
	if result.Error != nil {
		return nil, result.Error
	}
	return textList, nil
}

func GetTextAllList() ([]models.Textdata, error) {
	var textdata models.Textdata
	result := database.DB2.Find(&textdata)
	if result.Error != nil {
		return nil, result.Error
	}
	var textList []models.Textdata
	result = database.DB2.Find(&textList)
	if result.Error != nil {
		return nil, result.Error
	}
	return textList, nil
}

func UpdateTextData(text models.Textdata) error {
	result := database.DB2.Save(&text)
	return result.Error
}
