package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hello/core"
	"github.com/hello/services"
	"github.com/hello/typings"
	validatorsTypings "github.com/hello/validators/typings"
)

func GetEvents(httpContext *gin.Context) {
	query := httpContext.MustGet("reqQuery").(*validatorsTypings.GetEventsRequestQueryType)
	fmt.Println("[test] reqQuery ->", query, query.PageNo)
	result, err := services.GetEvents(query.PageNo, query.PageSize)
	httpContext.JSON(core.ApiResponse(result, err))
}

func CreateEvent(httpContext *gin.Context) {
	body := httpContext.MustGet("reqBody").(*validatorsTypings.CreateEventRequestBodyType)
	fmt.Println("[test] reqBody ->", body)
	var payload typings.CreateEventPayloadType
	payload.Name = body.Name
	payload.Description = body.Description
	payload.CreatedAt = time.Now()
	payload.CreatedAt = time.Now()

	services.CreateEvent(&payload)
	fmt.Println(body)
}
