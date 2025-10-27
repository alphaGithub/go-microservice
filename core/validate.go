package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type validatorFuncType func() error

func ValidateRequest(validatorFunc validatorFuncType) gin.HandlerFunc {
	fmt.Print(validatorFunc)
	return func(context *gin.Context) {
		// err := validatorFunc(context)
		// if err != nil {
		// 	context.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		// 	return
		// }
		context.Next()
	}
}
