package textcontroller

import (
	"SayLoveWall/app/models"
	textservice "SayLoveWall/app/services/textService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
)

type TextData struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" binding:"required"`
	Text   string `json:"text" binding:"required"`
	Unname string `json:"unname" binding:"required"`
}

func PublishOrNew(c *gin.Context) {

	var data TextData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	if data.ID == 0 {
		//判断字数
		if !textservice.WordCountJudge(data.Text) {
			utils.JsonErrorResponse(c, 200503, "字数过多")
			return
		}
		//存储
		err = textservice.Publish(models.Textdata{
			Name:   data.Name,
			Text:   data.Text,
			Unname: data.Unname,
		})
		if err != nil {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
		utils.JsonSuccessResponse(c, nil)
	} else {
		err = textservice.UpdateTextData(models.Textdata{
			ID:     data.ID,
			Name:   data.Name,
			Text:   data.Text,
			Unname: data.Unname,
		})
		if err != nil {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
		utils.JsonSuccessResponse(c, nil)
	}

}
