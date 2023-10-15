package usercontroller

import (
	"SayLoveWall/app/models"
	userservice "SayLoveWall/app/services/userService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
)

type NewUserData struct {
	ID       uint   `json:"id" binding:"required"`
	PhoneNum string `json:"phone_num" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Sex      int    `json:"gender" binding:"required"`
	Age      uint   `json:"age" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UpdateUserData(c *gin.Context) {
	//接收参数
	var data NewUserData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断手机号位数
	if len(data.PhoneNum) != 11 {
		utils.JsonErrorResponse(c, 200502, "手机号格式错误")
		return
	}

	err = userservice.UpdateUserData(models.Userdata{
		ID:       data.ID,
		Name:     data.Name,
		Sex:      data.Sex,
		PhoneNum: data.PhoneNum,
		Age:      data.Age,
		Password: utils.Encrypt(data.Password),
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
