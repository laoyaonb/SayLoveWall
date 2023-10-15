package usercontroller

import (
	"SayLoveWall/app/models"
	userservice "SayLoveWall/app/services/userService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	PhoneNum        string `json:"phone_num" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Sex             string `json:"sex" binding:"required"`
	Agree           string `json:"agree" binding:"required"`
}

func Register(c *gin.Context) {
	//接收参数
	var data RegisterData
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

	//判断手机号是否已被注册
	err = userservice.CheckUserExistByPhoneNum(data.PhoneNum)
	if err == nil {
		utils.JsonErrorResponse(c, 200503, "手机号已被注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//判断密码是否一致
	flag := userservice.ComparePwd(data.Password, data.ConfirmPassword)
	if !flag {
		utils.JsonErrorResponse(c, 200504, "密码不一致")
		return
	}

	//判断是否同意协议
	if data.Agree != "1" {
		utils.JsonErrorResponse(c, 200505, "未能同意协议")
		return
	}

	//用户注册
	err = userservice.Register(models.Userdata{
		Name:     data.Name,
		Sex:      data.Sex,
		PhoneNum: data.PhoneNum,
		Password: data.Password,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)

}
