package model

// Hello 的参数信息
//
// swagger:parameters HelloParam
type HelloParam struct {
	// Name
	// 例如: world
	//
	// Required: true
	// In: query
	Name string `json:"name"`
}
