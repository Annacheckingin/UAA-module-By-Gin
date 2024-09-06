package main

import (
	middlewares "UAAByGin/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()

	middlewares.SetUp()
	gin.Run(":8080")
}
