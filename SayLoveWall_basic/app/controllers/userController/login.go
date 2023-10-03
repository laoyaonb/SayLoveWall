package usercontroller

import (
	"SayLoveWall/app/models"
	"SayLoveWall/app/services/userService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	PhoneNum string `json:"phone_num" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

//登录
func Login(c *gin.Context) {
	//接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断用户账号是否存在
	err = userservice.CheckUserExistByPhoneNum(data.PhoneNum)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	//获取用户信息
	var userdata *models.Userdata
	userdata, err = userservice.GetUserByPhoneNum(data.PhoneNum)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//判断密码是否正确

	flag := userservice.ComparePwd(data.PassWord, userdata.Password)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "密码错误")
		return
	}

	utils.JsonSuccessResponse(c, userdata)

}
