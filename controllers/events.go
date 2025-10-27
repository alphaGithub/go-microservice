package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hello/validators/typings"
)

func GetEvents(context *gin.Context) {
	query := context.MustGet("reqQuery").(*typings.GetEventsRequestQueryType)
	fmt.Print("reqQuery ->", query, query.PageNo)
	context.JSON(http.StatusOK, gin.H{"success": true})
}
