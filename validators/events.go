package validators

type GetEventsRequestQuery struct {
	PageNo   int `json:"page_no" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}
