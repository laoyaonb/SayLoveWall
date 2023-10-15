package router

import (
	usercontroller "SayLoveWall/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/login", usercontroller.Login)
		api.POST("/register", usercontroller.Register)
	}

}
