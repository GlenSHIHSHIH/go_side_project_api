package dto

type BaseResponseDTO struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
