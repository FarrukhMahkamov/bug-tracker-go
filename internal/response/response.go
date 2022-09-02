package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

type NewSuccessResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
