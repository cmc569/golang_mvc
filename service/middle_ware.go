package service

import (
	_ "api/util/jwt"
	util "api/util/line_webhook"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var isPass = true
		var message string
		token := c.Request.Header.Get("authorization")
		if token == "" {
			message = "token not found"
			isPass = false
		} else {
			rsp := util.VerifyAccessToken(token)
			if rsp.UserId == ""{
				message = rsp.Message
				isPass = false
			}
		}
		if !isPass {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": message,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
