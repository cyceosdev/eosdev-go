package controller

import "eosdev-go/code/model"

func HelloCtrl(param *model.HelloParam) (data *model.HelloResponse, err *model.ErrorResponse) {
	data = &model.HelloResponse{}

	if param.Name != "" {
		data.Res = "Hello " + param.Name
		return data, nil
	} else {
		return nil, &model.ErrorResponse{Code: 1000, Describe: "name 参数是一个空值"}
	}
}
