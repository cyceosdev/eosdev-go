package model

type ErrorResponse struct {
	Code     int64  `json:"code"`
	Describe string `json:"describe"`
}

var ErrNameNull = ErrorResponse{Code: 1000, Describe: "name 参数是一个空值"}
