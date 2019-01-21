package controller

import "eosdev-go/code/model"

func HelloCtrl(params *model.HelloParams) (data *model.HelloResponse, err *model.ErrorResponse) {
	data = &model.HelloResponse{}

	if params.Name != "" {
		data.Res = "Hello " + params.Name
		return data, nil
	} else {
		return nil, &model.ErrNameNull
	}
}
