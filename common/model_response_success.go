package common

type success_response struct {
	Paging  interface{} `json:"paging,omitempty"`
	Filter  interface{} `json:"filter,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SuccessResponse_Full(paging, filter, data interface{}) *success_response {
	return &success_response{
		Paging: paging,
		Filter: filter,
		Data:   data,
	}
}

func SuccessResponse_Data(data interface{}) *success_response {
	return &success_response{
		Paging: nil,
		Filter: nil,
		Data:   data,
	}
}

func SuccessResponse_Message(mess string) *success_response {
	return &success_response{
		Paging:  nil,
		Filter:  nil,
		Data:    nil,
		Message: mess,
	}
}
