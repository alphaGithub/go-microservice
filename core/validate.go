package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequest(reqStruct any) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := context.BindQuery(reqStruct)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		context.Next()
	}
}
