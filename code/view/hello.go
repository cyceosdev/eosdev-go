package view

import (
	"eosdev-go/code/controller"
	"eosdev-go/code/lib"
	"eosdev-go/code/model"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	// swagger:route GET /hello 发币 HelloParams
	//
	// 打印Hello
	//
	// 打印类似 “Hello world” 的语句
	//
	// Responses:
	//   200: HelloResponse
	//   500: ErrorResponse

	var queryVals = req.URL.Query()
	var name = queryVals.Get("name")

	var params = &model.HelloParams{Name: name}

	if data, err := controller.HelloCtrl(params); err == nil {
		w.WriteHeader(200)
		w.Write(lib.StructToJson(data))
	} else {
		w.WriteHeader(500)
		w.Write(lib.StructToJson(err))
	}
}
