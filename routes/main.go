package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(user *gin.RouterGroup) {
	user.POST("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "login!",
			"success": "true",
		})
	})
}

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": "true",
			"message": "ok.",
		})
	})
	userRoutes := server.Group("/user")

	registerUserRoutes(userRoutes)
}
