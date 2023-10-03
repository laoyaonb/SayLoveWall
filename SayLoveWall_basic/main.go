package main

import (
	"SayLoveWall/config/database"
	"SayLoveWall/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	database.Init()

	r := gin.Default()
	router.Init(r)

	err := r.Run()
	if err != nil{
		log.Fatal("Server start error: ", err)
	}
}