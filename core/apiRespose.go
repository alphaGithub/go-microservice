package core

import (
	"fmt"
	"net/http"
)

type ApiResponseOptions struct {
	Message string
	Status  int
}

type ResponseOptionFunc func(*ApiResponseOptions)

func WithMessage(message string) ResponseOptionFunc {
	return func(options *ApiResponseOptions) {
		options.Message = message
	}
}

func WithStatus(status int) ResponseOptionFunc {
	return func(options *ApiResponseOptions) {
		options.Status = status
	}
}

func ApiResponse(data interface{}, err error, opts ...ResponseOptionFunc) (int, map[string]interface{}) {
	fmt.Println("[test] >>>>>", data, err, opts)
	options := ApiResponseOptions{
		Message: "",
		Status:  0,
	}

	if len(opts) > 0 {
		opts[0](&options)
	}
	var response = make(map[string]interface{})
	if err != nil {
		response["success"] = false
		response["message"] = "Oops! something went wrong."
		response["data"] = nil
		response["error"] = err.Error()
		response["status"] = http.StatusBadRequest
	} else {
		response["success"] = true
		response["message"] = "success"
		response["data"] = data
		response["error"] = nil
		response["status"] = http.StatusOK
	}
	if len(options.Message) > 0 {
		response["message"] = options.Message
	}
	if options.Status > 0 {
		response["status"] = options.Status
	}
	statusValue, _ := response["status"].(int)
	fmt.Println("api response", statusValue, response)
	return statusValue, response
}
