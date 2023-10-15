package main

import (
	"SayLoveWall/app/midwares"
	"SayLoveWall/config/database"
	"SayLoveWall/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"}) //47.96.100.92
	r.Use(midwares.Cors())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
