package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequestQuery(reqStruct any) gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		err := httpContext.ShouldBindQuery(reqStruct)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		httpContext.Set("reqQuery", reqStruct)
		httpContext.Next()
	}
}

func ValidateRequestBody(reqStruct any) gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		err := httpContext.ShouldBindJSON(reqStruct)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		httpContext.Set("reqBody", reqStruct)
		httpContext.Next()
	}
}

func ValidateRequestParams(reqStruct any) gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		err := httpContext.ShouldBindUri(reqStruct)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		httpContext.Set("reqParams", reqStruct)
		httpContext.Next()
	}
}
