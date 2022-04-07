package v1

import (
	"api/util/gin_extend"
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	c := gin_extend.BeforeAction(context)
	c.SuccessResponse("success")
}
