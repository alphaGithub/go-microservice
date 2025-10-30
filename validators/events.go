package validators

import "github.com/hello/validators/typings"

func GetEventsRequestQuery() *typings.GetEventsRequestQueryType {
	return &typings.GetEventsRequestQueryType{PageNo: 1, PageSize: 10}
}

func CreateEventRequestBody() *typings.CreateEventRequestBodyType {
	return &typings.CreateEventRequestBodyType{}
}

func GetEventByIdRequestParams() *typings.GetEventByIdRequestParamsType {
	return &typings.GetEventByIdRequestParamsType{}
}
