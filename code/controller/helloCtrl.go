package controller

import "eosdev-go/code/model"

func HelloCtrl(param *model.HelloParams) (data *model.HelloResponse, err *model.ErrorResponse) {
	data = &model.HelloResponse{}

	if param.Name != "" {
		data.Res = "Hello " + param.Name
		return data, nil
	} else {
		return nil, &model.ErrNameNull
	}
}
