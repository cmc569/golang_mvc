package gin_extend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	Context *gin.Context
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BeforeAction(content *gin.Context) *Gin {
	c := Gin{Context: content}
	return &c
}

func (gin *Gin) SuccessResponse(msg string, data ...interface{}) {
	var resData interface{}
	if data == nil {
		resData = ""
	} else {
		resData = data[0]
	}
	gin.Context.JSON(http.StatusOK, Response{
		Msg:  msg,
		Data: resData,
	})
	return
}

func (gin *Gin) ErrorResponse(msg string, data ...interface{}) {
	var resData interface{}
	if data == nil {
		resData = ""
	} else {
		resData = data[0]
	}

	gin.Context.JSON(http.StatusBadRequest, Response{
		Msg:  msg,
		Data: resData,
	})
	return
}
