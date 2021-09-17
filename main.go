package main

import (

	"github.com/gin-gonic/gin"
	"github.com/luvkrai/mytestablecodes/controllers"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.PingController.Ping)

	router.Run(":8080")
}