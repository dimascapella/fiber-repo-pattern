package helper

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func ErrorResponse(message string, err string, data interface{}) Response {
	res := Response{
		Status:  false,
		Message: message,
		Errors:  err,
		Data:    data,
	}
	return res
}

type EmptyObj struct{}
