package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hello/validators"
)

func GetEvents(context *gin.Context) {
	req := context.MustGet("reqQuery").(*validators.GetEventsRequestQuery)
	fmt.Print("reqBody ->", req, req.PageNo)
	context.JSON(http.StatusOK, gin.H{"success": true})
}
