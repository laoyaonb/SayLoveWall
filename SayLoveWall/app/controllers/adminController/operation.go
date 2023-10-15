package admincontroller

import (
	"SayLoveWall/app/models"
	textservice "SayLoveWall/app/services/textService"
	"SayLoveWall/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

type TextData struct {
	ID     uint   `json:"id" binding:"required"`
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

	var a string
	if data.ID == 0 {
		a = "false"
	} else {
		a = "true"
	}
	utils.JsonErrorResponse(c, 200505, a)

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
