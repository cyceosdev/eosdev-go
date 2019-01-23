package controllers

import (
	"encoding/json"

	"com.zhaoyin/eosdev-go/models"
	"github.com/astaxie/beego"
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
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
		remp["result"] = err.Error()
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
	var query = c.Ctx.Input.Context.Request.URL.Query()

	var strAccountName = query.Get("name")
	var name = eos.AN(strAccountName)
	if out, err := models.GetAccount(name); err != nil {
		remp["result"] = err.Error()
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

func (c *RestController) IssueToken() {
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

	var strTo = request["to"]
	var strQuantity = request["quantity"]
	var memo = request["memo"]

	var to = eos.AN(strTo)

	var quantity eos.Asset
	if val, err := eos.NewAsset(strQuantity); err != nil {
		remp["result"] = err.Error()
		remp["state"] = 2
		c.ReturnValue(remp)
		return
	} else {
		quantity = val
	}

	if out, err := models.IssueToken(to, quantity, memo); err != nil {
		remp["result"] = err.Error()
		remp["state"] = 2
		c.ReturnValue(remp)
		return
	} else {
		remp["result"] = out
		remp["state"] = 0
		c.ReturnValue(remp)
		return
	}

}

func (c *RestController) GetCurrencyBalance() {
	// account eos.AccountName, symbol string, code eos.AccountName
	var remp = make(map[string]interface{})
	var query = c.Ctx.Input.Context.Request.URL.Query()

	var strAccountName = query.Get("name")
	var symbol = query.Get("symbol")
	var strCode = query.Get("code")

	var account = eos.AN(strAccountName)
	var code = eos.AN(strCode)

	if out, err := models.GetCurrencyBalance(account, symbol, code); err != nil {
		remp["result"] = err.Error()
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

func (c *RestController) CreateAccount() {
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

	var strName = request["name"]
	var strPublicKey = request["public_key"]

	var accountName = eos.AN(strName)
	var publicKey ecc.PublicKey

	if val, err := ecc.NewPublicKey(strPublicKey); err != nil {
		remp["result"] = err.Error()
		remp["state"] = 2
		c.ReturnValue(remp)
		return
	} else {
		publicKey = val
	}

	if out, err := models.CreateAccount(accountName, publicKey); err != nil {
		remp["result"] = err.Error()
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

func (c *RestController) RootTransfer() {
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

	var strTo = request["to"]
	var strQuantity = request["quantity"]
	var memo = request["memo"]

	var to = eos.AN(strTo)

	var quantity eos.Asset
	if val, err := eos.NewAsset(strQuantity); err != nil {
		remp["result"] = err.Error()
		remp["state"] = 2
		c.ReturnValue(remp)
		return
	} else {
		quantity = val
	}

	if out, err := models.RootTransfer(to, quantity, memo); err != nil {
		remp["result"] = err.Error()
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
