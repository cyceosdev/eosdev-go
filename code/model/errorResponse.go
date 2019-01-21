package model

type ErrorResponse struct {
	Code     int64  `json:"code"`
	Describe string `json:"describe"`
}

// 错误的响应信息
//
// swagger:response ErrorResponse
type ErrorResponseWrapper struct {
	// in: body
	Body ErrorResponse
}

var ErrNameNull = ErrorResponse{Code: 1000, Describe: "name 参数是一个空值"}
