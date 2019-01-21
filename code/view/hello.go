package view

import (
	"eosdev-go/code/controller"
	"eosdev-go/code/lib"
	"eosdev-go/code/model"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	var queryVals = req.URL.Query()
	var name = queryVals.Get("name")

	var param = &model.HelloParam{Name: name}

	if data, err := controller.HelloCtrl(param); err == nil {
		w.WriteHeader(200)
		w.Write(lib.StructToJson(data))
	} else {
		w.WriteHeader(500)
		w.Write(lib.StructToJson(err))
	}
}
