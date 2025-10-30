package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hello/controllers"
	"github.com/hello/core"
	"github.com/hello/validators"
)

func registerEventsRoutes(events *gin.RouterGroup) {
	events.GET("/", core.ValidateRequestQuery(validators.GetEventsRequestQuery()), controllers.GetEvents)
	events.GET("/:id", core.ValidateRequestParams(validators.GetEventByIdRequestParams()), controllers.GetEventById)
	events.POST("/", core.ValidateRequestBody(validators.CreateEventRequestBody()), controllers.CreateEvent)
	events.PUT("/:id", controllers.GetEvents)
	events.DELETE("/:id", controllers.GetEvents)
}
func registerUserRoutes(user *gin.RouterGroup) {
	user.POST("/login", func(httpContext *gin.Context) {
		httpContext.JSON(http.StatusOK, gin.H{
			"message": "login!",
			"success": "true",
		})
	})
}

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", func(httpContext *gin.Context) {
		httpContext.JSON(http.StatusOK, gin.H{
			"success": "true",
			"message": "ok.",
		})
	})
	userRoutes := server.Group("/user")
	eventsRoutes := server.Group("/events")

	registerUserRoutes(userRoutes)
	registerEventsRoutes(eventsRoutes)
}
