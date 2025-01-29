package dto

type ResponseDTO struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}
