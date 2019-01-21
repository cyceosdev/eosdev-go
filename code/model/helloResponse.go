package model

type HelloResponse struct {
	Res string `json:"res"`
}

// Hello 的响应信息
//
// swagger:response HelloResponse
type HelloResponseWrapper struct {
	// in: body
	Body HelloResponse
}
