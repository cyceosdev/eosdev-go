package controllers

import (
	"encoding/json"

	"com.zhaoyin/eosdev-go/models"
	"github.com/astaxie/beego"
	eos "github.com/eoscanada/eos-go"
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

func (c *RestController) CreateToken() {
	var remp = make(map[string]interface{})
	var request map[string]string
	if c.Ctx.Input.RequestBody == nil {
		c.ReturnValue(remp)
		return
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		remp["result"] = err.Error()
		remp["state"] = 1
		c.ReturnValue(remp)
		return
	}

	var strMaxSupply = request["max_supply"]

	var maxSupply eos.Asset
	if val, err := eos.NewAsset(strMaxSupply); err != nil {
		remp["result"] = err.Error()
		remp["state"] = 2
		c.ReturnValue(remp)
		return
	} else {
		maxSupply = val
	}

	if out, err := models.CreateToken(maxSupply); err != nil {
		remp["result"] = err
		remp["state"] = 3
		c.ReturnValue(remp)
		return
	} else {
		remp["result"] = out
		remp["state"] = 0
		c.ReturnValue(remp)
		return
	}
}

func (c *RestController) GetAccount() {
	var remp = make(map[string]interface{})
	if c.Ctx.Input.RequestBody == nil {
		c.ReturnValue(remp)
		return
	}

	var query = c.Ctx.Input.Context.Request.URL.Query()

	var strAccountName = query.Get("name")
	var name = eos.AN(strAccountName)
	if out, err := models.GetAccount(name); err != nil {
		remp["result"] = err
		remp["state"] = 3
		c.ReturnValue(remp)
		return
	} else {
		remp["result"] = out
		remp["state"] = 0
		c.ReturnValue(remp)
		return
	}
}

func (c *RestController) ReturnValue(remp map[string]interface{}) {
	c.Data["json"] = remp
	c.ServeJSON()
}
