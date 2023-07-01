package app

import (
	"github.com/cindyyangcaixia/gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(message *ResponseMessage, data interface{}) {
	if message == nil {
		g.C.JSON(e.SUCCESS, Response{
			Code:    e.SUCCESS,
			Message: e.GetMsg(e.SUCCESS),
			Data:    data,
		})
	} else {
		msg := e.GetMsg(message.Code)
		if message.Message != "" {
			msg = message.Message
		}
		g.C.JSON(message.Status, Response{
			Code:    message.Code,
			Message: msg,
			Data:    data,
		})
	}
}
