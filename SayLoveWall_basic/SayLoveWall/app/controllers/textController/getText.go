package textcontroller

import (
	"SayLoveWall/app/models"
	textservice "SayLoveWall/app/services/textService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTextPersonalData struct {
	Name string `form:"name" binding:"required"`
}

func GetTextPersonal(c *gin.Context) {
	var data GetTextPersonalData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	var textList []models.Textdata
	textList, err = textservice.GetTextPersonalList(data.Name)
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
