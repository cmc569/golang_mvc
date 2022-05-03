package main

import (
	v1 "api/api/v1"
	_ "api/docs"
	"api/service"
	"net/http"

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
	// router.GET('/login')
}

func Web(router *gin.RouterGroup) {
	router.GET("/:name/:user_id", showUserInfo)
	router.GET("/", showAction)
	router.POST("/posts", showPosts)
}

func showUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": c.Param("name"),
		"userId": c.Param("user_id"),
	})
}

func showAction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"action": c.Query("method"),
		"data": c.Query("data"),
	})
}
/*
raw json:
{
    "post_id": 12,
    "title": "Test Text"
}
*/
func showPosts(c *gin.Context) {
	var response map[string]interface{}
	err := c.Bind(&response)
	if err != nil {
		return;
	}

	c.JSON(200, gin.H{
		"id": response["post_id"],
	})
	// or 
	// c.JSON(200, gin.H(response))
}
