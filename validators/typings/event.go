package typings

import "time"

type GetEventsRequestQueryType struct {
	PageNo   int `form:"page_no"`
	PageSize int `form:"page_size"`
}

type CreateEventRequestBodyType struct {
	Name        string                 `bindings:"required" json:"name"`
	Description string                 `bindings:"required" json:"description"`
	Payload     map[string]interface{} `bindings:"required" json:"payload"`
}

type CreateEventPayloadType struct {
	Name        string
	Description string
	Payload     map[string]interface{}
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
