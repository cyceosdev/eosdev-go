package model

type ErrorResponse struct {
	Code     int64  `json:"code"`
	Describe string `json:"describe"`
}
