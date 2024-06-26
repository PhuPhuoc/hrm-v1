package common

import "net/http"

type success_response struct {
	Status  int         `json:"status"`
	Paging  interface{} `json:"paging,omitempty"`
	Filter  interface{} `json:"filter,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SuccessResponse_Full(status int, paging, filter, data interface{}) *success_response {
	return &success_response{
		Status: status,
		Paging: paging,
		Filter: filter,
		Data:   data,
	}
}

func SuccessResponse_GetObject(paging, filter, data interface{}) *success_response {
	return &success_response{
		Status: http.StatusOK,
		Paging: paging,
		Filter: filter,
		Data:   data,
	}
}

func SuccessResponse_Data(data interface{}) *success_response {
	return &success_response{
		Status: http.StatusOK,
		Paging: nil,
		Filter: nil,
		Data:   data,
	}
}

func SuccessResponse_Message(mess string) *success_response {
	return &success_response{
		Status:  http.StatusCreated,
		Paging:  nil,
		Filter:  nil,
		Data:    nil,
		Message: mess,
	}
}

func SuccessResponse_NoContent(mess string) *success_response {
	return &success_response{
		Status:  http.StatusNoContent,
		Paging:  nil,
		Filter:  nil,
		Data:    nil,
		Message: mess,
	}
}
