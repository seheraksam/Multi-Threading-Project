package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seheraksam/Multi-Threading-Project/controllers"
	"github.com/seheraksam/Multi-Threading-Project/initializers"
)

func init() {
	initializers.ConnectRedis()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/scrape", controllers.GetTitle)

	r.Run(":8080")
}
