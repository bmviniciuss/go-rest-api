package http_helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorResponse(c *gin.Context, statusCode int, msg string) {
	c.JSON(statusCode, gin.H{
		"message": msg,
	})
}

func successResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func BadRequestErrorResponse(c *gin.Context, msg string) {
	errorResponse(c, http.StatusBadRequest, msg)
}

func ServerErrorResponse(c *gin.Context, msg string) {
	errorResponse(c, http.StatusInternalServerError, msg)
}

func NotFoundErrorResponse(c *gin.Context, msg string) {
	errorResponse(c, http.StatusNotFound, msg)
}

func OkResponse(c *gin.Context, data interface{}) {
	successResponse(c, http.StatusOK, data)
}
