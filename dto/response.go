package dto

type Response struct {
	HttpStatus int         `json:"httpStatus,omitempty"`
	Message    string      `json:"message,omitempty"`
	Response   interface{} `json:"response,omitempty"`
}
