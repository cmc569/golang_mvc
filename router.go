package main

import (
	v1 "api/api/v1"
	_ "api/docs"
	"api/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Content-Type", "Access-Control-Allow-Origin", "Authorization"},
	}))

	webGroup := router.Group("/web")
	Web(webGroup.Group("/"))
	//======================================================================
	//								  not auth
	//======================================================================
	apiV1Group := router.Group("/api/v1")
	Public(apiV1Group.Group("/public"))

	//======================================================================
	//								  auth
	//======================================================================
	apiV1AuthGroup := router.Group("/api/v1/auth")
	apiV1AuthGroup.Use(service.AuthMiddleware())
	{
	}
	return router
}

func Public(router *gin.RouterGroup) {
	router.POST("/test", v1.Ping)
}

func Web(router *gin.RouterGroup) {
	router.GET("/hello", func (c *gin.Context) {
		c.JSON(200, "Hello - v1")
	})
}