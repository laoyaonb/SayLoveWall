package textcontroller

import (
	"SayLoveWall/app/models"
	textservice "SayLoveWall/app/services/textService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTextPersonal(c *gin.Context) {
	Name := c.Param("username")

	var textList []models.Textdata
	var err error
	textList, err = textservice.GetTextPersonalList(Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "表白列表为空")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	utils.JsonSuccessResponse(c, gin.H{
		"text_list": textList,
		"num":       len(textList),
	})
}

func GetTextAll(c *gin.Context) {
	var textList []models.Textdata
	var err error
	textList, err = textservice.GetTextAllList()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "表白列表为空")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	utils.JsonSuccessResponse(c, gin.H{
		"text_list": textList,
		"num":       len(textList),
	})
}
