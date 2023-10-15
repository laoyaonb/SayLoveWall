package textcontroller

import (
	textservice "SayLoveWall/app/services/textService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteData struct {
	ID uint `json:"id"`
}

func Delete(c *gin.Context) {
	//接收参数
	var data DeleteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//删除
	err = textservice.Delete(data.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "未找到该记录")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
}
