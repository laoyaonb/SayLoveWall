package admincontroller

import (
	"SayLoveWall/app/models"
	adminservice "SayLoveWall/app/services/adminService"
	"SayLoveWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Account  string `json:"Account" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	//接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断管理员账号是否存在
	err = adminservice.CheckUserExistByAccount(data.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "管理员账号不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	//获取用户信息
	var admindata *models.Admindata
	admindata, err = adminservice.GetAdminByAccount(data.Account)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//判断密码是否正确

	flag := adminservice.ComparePwd(data.PassWord, admindata.Password)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "密码错误")
		return
	}

	utils.JsonSuccessResponse(c, admindata)
}
