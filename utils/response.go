package util

import "go-clean-code/constant"

type BodyResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func response(success bool, message string, data interface{}, err any) BodyResponse {
	return BodyResponse{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

func ResponseSuccess(message string, data interface{}) BodyResponse {
	return response(true, message, data, nil)
}

func ResponseNotFound(message string) BodyResponse {
	var newMessage string
	if message == "" {
		newMessage = constant.NotFound()
	} else {
		newMessage = message
	}
	return response(false, newMessage, nil, nil)
}

func ResponseError(message string, error any) BodyResponse {
	return response(false, message, nil, error)
}
