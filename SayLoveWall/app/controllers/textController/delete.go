package textcontroller

import (
	textservice "SayLoveWall/app/services/textService"
	"SayLoveWall/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(c *gin.Context) {
	//接收参数
	sID := c.Param("id")
	iID, _ := strconv.Atoi(sID)
	//删除
	err := textservice.Delete(uint(iID))
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
