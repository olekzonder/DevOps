package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.POST("/project_name", BuildProject)
	router.Run(":9090")
}
