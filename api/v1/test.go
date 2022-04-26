package v1

import (
	"api/util/gin_extend"
	"github.com/gin-gonic/gin"
	"time"
)

func Ping(context *gin.Context) {
	c := gin_extend.BeforeAction(context)
	time.Sleep(3 * time.Second)
	c.SuccessResponse("success")

}
