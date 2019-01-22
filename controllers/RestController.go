package controllers

import (
	"com.zhaoyin/etoken-go/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type RestController struct {
	beego.Controller
}

func (c *RestController) OpenAccount() {
	remp := make(map[string]interface{})
	var request map[string]string
	if c.Ctx.Input.RequestBody != nil {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
		if err != nil {
			remp["result"] = err.Error()
			remp["state"] = 1
		} else {
			remp["result"] = "开户请求成功!"
			remp["state"] = 0
		}
	}
	c.Data["json"] = remp
	c.ServeJSON()
}

func (c *RestController) TestApi() {
	remp := make(map[string]interface{})
	var request map[string]string
	if c.Ctx.Input.RequestBody != nil {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
		if err != nil {
			remp["result"] = err.Error()
			remp["state"] = 1
		} else {
			out, err := models.GetEosBalance()
			info, _ := models.GetInfo()
			if err != nil {
				remp["result"] = err.Error()
				remp["state"] = 1
			} else {
				remp["result"] = out
				remp["info"] = info
				remp["state"] = 0
			}
		}
	}
	c.Data["json"] = remp
	c.ServeJSON()
}
