package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequestQuery(reqStruct any) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := context.ShouldBindQuery(reqStruct)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		context.Set("reqQuery", reqStruct)
		context.Next()
	}
}

func ValidateRequestBody(reqStruct any) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := context.ShouldBindJSON(reqStruct)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		context.Set("reqBody", reqStruct)
		context.Next()
	}
}

func ValidateRequestParams(reqStruct any) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := context.ShouldBindUri(reqStruct)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		context.Set("reqParams", reqStruct)
		context.Next()
	}
}
