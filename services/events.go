package services

import (
	"fmt"

	"github.com/hello/dao"
	"github.com/hello/models"
	"github.com/hello/typings"
)

func GetEvents(pageNo int, pageSize int) ([]models.Event, error) {
	fmt.Println("[msg] service call")
	offset := (pageNo - 1) * pageSize
	limit := pageSize
	result, err := dao.GetEvents(limit, offset)
	if result == nil {
		result = []models.Event{}
	}
	return result, err
}

func GetEventById(id string) (models.Event, error) {
	fmt.Println("[msg] service call")
	result, err := dao.GetEventById(id) // sort_id
	return result, err
}

func CreateEvent(payload *typings.CreateEventPayloadType) (interface{}, error) {
	var event models.Event
	event.ShortId = payload.ShortId
	event.Payload = payload.Payload
	event.Name = payload.Name
	event.Description = payload.Description
	event.CreatedAt = payload.CreatedAt
	event.UpdatedAt = payload.UpdatedAt
	result, err := dao.CreateEvent(&event)
	if result != nil {
		var response = make(map[string]interface{})
		response["id"] = result
		return response, err
	}

	return result, err
}
