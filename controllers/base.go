package controllers

type Response struct {
	Data    interface{}
	Err     interface{}
	Success bool
	Message string
	ResCode string
}

func SuccessResponse(message string, data interface{}) *Response {
	return &Response{
		Data:    data,
		Success: true,
		Message: message,
		ResCode: "200",
	}
}

func FailResponse(message string, err interface{}) *Response {
	return &Response{
		Err:     err,
		Success: false,
		Message: message,
		ResCode: "500",
	}
}
