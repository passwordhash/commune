package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, status int, msg string) {
	logrus.Error(msg)
	c.AbortWithStatusJSON(status, response{msg})
}
