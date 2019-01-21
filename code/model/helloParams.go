package model

// Hello 的参数信息
//
// swagger:parameters HelloParams
type HelloParams struct {
	// Name
	// 例如: world
	//
	// Required: true
	// In: query
	Name string `json:"name"`
}
