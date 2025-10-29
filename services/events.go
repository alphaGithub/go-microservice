package services

import (
	"fmt"

	"github.com/hello/dao"
	"github.com/hello/models"
	"github.com/hello/typings"
)

func GetEvents(pageNo int, pageSize int) ([]models.Events, error) {
	fmt.Println("[msg] service call")
	offset := (pageNo - 1) * pageSize
	limit := pageSize
	result, err := dao.GetEvents(limit, offset)
	if result == nil {
		result = []models.Events{}
	}
	return result, err
}

func CreateEvent(payload *typings.CreateEventPayloadType) {

}
