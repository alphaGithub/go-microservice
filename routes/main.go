package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hello/core"
	"github.com/hello/validators"
)

func registerEventsRoutes(events *gin.RouterGroup) {
	events.GET("/list", core.ValidateRequest(validators.GetEventsRequestQuery), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "events list!",
			"success": "true",
		})
	})
}
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
	eventsRoutes := server.Group("/events")

	registerUserRoutes(userRoutes)
	registerEventsRoutes(eventsRoutes)
}
