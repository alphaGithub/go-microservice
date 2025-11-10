package typings

type GetEventsRequestQueryType struct {
	PageNo   int `form:"page_no"`
	PageSize int `form:"page_size"`
}

type CreateEventRequestBodyType struct {
	Name        string                 `bindings:"required" json:"name"`
	Description string                 `bindings:"required" json:"description"`
	MetaData    map[string]interface{} `bindings:"required" json:"meta_data"`
}

type GetEventByIdRequestParamsType struct {
	Id string `bindings:"required" uri:"id"`
}
