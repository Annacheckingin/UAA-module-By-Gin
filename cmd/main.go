package main

import (
	"UAAByGin/config"
	_ "UAAByGin/config"
	"UAAByGin/controllers"
	_ "UAAByGin/controllers"
	"UAAByGin/middlewares"
	_ "UAAByGin/middlewares"
	"UAAByGin/models"
	_ "UAAByGin/models"
	_ "UAAByGin/repositories"
	"UAAByGin/routers"
	_ "UAAByGin/routers"
	"UAAByGin/services"
	_ "UAAByGin/services"
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	config.SetUp(gin)
	middlewares.SetUp(gin)
	models.SetUp(gin)
	controllers.SetUp(gin)
	services.SetUp(gin)
	routers.SetUp(gin)
	gin.Run(":8080")
}
