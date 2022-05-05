package main

import (
	"redispractice/src"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "redispractice/docs"
)

// @title API for data in Redis
// @version 1.0
// @description Gin swagger (to CRUD User Info)

// @contact.name Jessie Shen

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http
func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	src.AddUserRouter(v1)
	router.Run(":8080")
}
