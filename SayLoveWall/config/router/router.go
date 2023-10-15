package router

import (
	admincontroller "SayLoveWall/app/controllers/adminController"
	textcontroller "SayLoveWall/app/controllers/textController"
	usercontroller "SayLoveWall/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/login", usercontroller.Login)
		api.POST("/register", usercontroller.Register)
		api.POST("/publish", textcontroller.PublishOrNew)
		api.GET("/personal/:username", textcontroller.GetTextPersonal)
		api.GET("/public", textcontroller.GetTextAll)
		api.DELETE("/personal/:id", textcontroller.Delete)
		api.DELETE("/public/:id", textcontroller.Delete)

		api.POST("adminlogin", admincontroller.Login)
		api.GET("/admin", admincontroller.GetTextAll)
		api.POST("/admin", admincontroller.PublishOrNew)
		api.DELETE("/admin/:id", admincontroller.Delete)
	}

}
