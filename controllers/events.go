package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hello/core"
	"github.com/hello/services"
	"github.com/hello/typings"
	validatorsTypings "github.com/hello/validators/typings"
	terisIoShortId "github.com/teris-io/shortid"
)

func GetEvents(httpContext *gin.Context) {
	query := httpContext.MustGet("reqQuery").(*validatorsTypings.GetEventsRequestQueryType)
	fmt.Println("[test] reqQuery ->", query, query.PageNo)
	result, err := services.GetEvents(query.PageNo, query.PageSize)
	httpContext.JSON(core.ApiResponse(result, err))
}

func GetEventById(httpContext *gin.Context) {
	params := httpContext.MustGet("reqParams").(*validatorsTypings.GetEventByIdRequestParamsType)
	fmt.Println("[test] reqParams ->", params)
	result, err := services.GetEventById(params.Id)
	httpContext.JSON(core.ApiResponse(result, err))
}

func CreateEvent(httpContext *gin.Context) {
	body := httpContext.MustGet("reqBody").(*validatorsTypings.CreateEventRequestBodyType)
	fmt.Println("[test] reqBody ->", body)
	ShortID, _ := terisIoShortId.Generate()
	var payload typings.CreateEventPayloadType
	payload.Name = body.Name
	payload.Description = body.Description
	payload.ShortId = ShortID
	payload.Payload = body.Payload
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	result, err := services.CreateEvent(&payload)
	httpContext.JSON(core.ApiResponse(result, err))

}
