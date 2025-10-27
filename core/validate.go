package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequestQuery(reqStruct any) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := context.BindQuery(reqStruct)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		fmt.Print("reqStruct ->", reqStruct)
		context.Set("reqQuery", reqStruct)
		context.Next()
	}
}
