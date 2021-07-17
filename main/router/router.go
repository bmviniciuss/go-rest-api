package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigRoutes(router *gin.Engine, db *gorm.DB) {
	baseRouter := router.Group("api/v1")
	// /books
	CreateBooksRouter(baseRouter, db)
}
