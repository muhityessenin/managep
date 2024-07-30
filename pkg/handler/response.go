package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Status  int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
func newResponse(statusCode int, message string, data interface{}) Response {
	return Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
}
