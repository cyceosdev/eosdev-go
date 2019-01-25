package controllers

import (
	"com.zhaoyin/eosdev-go/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/eoscanada/eos-go"
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
	var responseData = &models.ResponseData{}
	defer c.ReturnResponsData(responseData)

	var request map[string]string
	if c.Ctx.Input.RequestBody == nil {
		responseData = models.ErrorBodyNil
		return
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		responseData.Error = err.Error()
		responseData.State = models.ErrorJsonState
		return
	}

	var strMaxSupply = request["max_supply"]

	var maxSupply eos.Asset
	if val, err := eos.NewAsset(strMaxSupply); err != nil {
		responseData.State = models.ErrorAssetState
		responseData.Error = err.Error()
		return
	} else {
		maxSupply = val
	}

	if out, err := models.CreateToken(maxSupply); err != nil {
		responseData.State = models.ErrorEOSAPIState
		responseData.Error = err.Error()
		return
	} else {
		responseData.State = models.SuccessState
		responseData.Result = out
		return
	}
}

func (c *RestController) GetAccount() {
	var responseData = &models.ResponseData{}
	defer c.ReturnResponsData(responseData)

	var query = c.Ctx.Input.Context.Request.URL.Query()

	var strAccountName = query.Get("name")
	var name = eos.AN(strAccountName)
	if out, err := models.GetAccount(name); err != nil {
		responseData.State = models.ErrorEOSAPIState
		responseData.Error = err.Error()
		return
	} else {
		responseData.State = models.SuccessState
		responseData.Result = out
		return
	}
}

func (c *RestController) IssueToken() {
	var responseData = &models.ResponseData{}
	defer c.ReturnResponsData(responseData)

	var request map[string]string
	if c.Ctx.Input.RequestBody == nil {
		responseData = models.ErrorBodyNil
		return
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		responseData.State = models.ErrorJsonState
		responseData.Error = err.Error()
		return
	}

	var strTo = request["to"]
	var strQuantity = request["quantity"]
	var memo = request["memo"]

	var to = eos.AN(strTo)

	var quantity eos.Asset
	if val, err := eos.NewAsset(strQuantity); err != nil {
		responseData.State = models.ErrorAssetState
		responseData.Error = err.Error()
		return
	} else {
		quantity = val
	}

	if out, err := models.IssueToken(to, quantity, memo); err != nil {
		responseData.State = models.ErrorEOSAPIState
		responseData.Error = err.Error()
		return
	} else {
		responseData.State = models.SuccessState
		responseData.Result = out
		return
	}

}

func (c *RestController) GetCurrencyBalance() {
	var responseData = &models.ResponseData{}
	defer c.ReturnResponsData(responseData)

	var query = c.Ctx.Input.Context.Request.URL.Query()

	var strAccountName = query.Get("name")
	var symbol = query.Get("symbol")
	var strCode = query.Get("code")

	var account = eos.AN(strAccountName)
	var code = eos.AN(strCode)

	if out, err := models.GetCurrencyBalance(account, symbol, code); err != nil {
		responseData.State = models.ErrorEOSAPIState
		responseData.Error = err.Error()
		return
	} else {
		responseData.State = models.SuccessState
		responseData.Result = out
		return
	}
}

//func NewAccountName() string {
//	const elems = "abcdefghijklmnopqrstuvwxyz12345"
//
//	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
//	var name string
//	name += string(elems[r.Intn(25)])
//	for i := 1; i < 12; i++ {
//		name += string(elems[r.Intn(30)])
//	}
//
//	return name
//}

func (c *RestController) CreateAccount() {
	var responseData = &models.ResponseData{}
	defer c.ReturnResponsData(responseData)
	var strAccountName = models.NewRandomAccount()
	var accountName = eos.AN(strAccountName)
	var randomPriKey *ecc.PrivateKey
	if val, err := ecc.NewRandomPrivateKey(); err != nil {
		responseData.State = models.ErrorPriKeyState
		responseData.Error = err.Error()
		return
	} else {
		randomPriKey = val
	}
	var publicKey = randomPriKey.PublicKey()

	if _, err := models.CreateAccount(accountName, publicKey); err != nil {
		responseData.State = models.ErrorEOSAPIState
		responseData.Error = err.Error()
		return
	} else {
		responseData.State = models.SuccessState
		responseData.Result = &models.NewAccountInfo{
			Name:   strAccountName,
			PubKey: randomPriKey.PublicKey().String(),
			PriKey: randomPriKey.String(),
		}
		return
	}
}

func (c *RestController) RootTransfer() {
	var responseData = &models.ResponseData{}
	defer c.ReturnResponsData(responseData)

	var request map[string]string
	if c.Ctx.Input.RequestBody == nil {
		responseData = models.ErrorBodyNil
		return
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		responseData.State = models.ErrorJsonState
		responseData.Error = err.Error()
		return
	}

	var strTo = request["to"]
	var strQuantity = request["quantity"]
	var memo = request["memo"]

	var to = eos.AN(strTo)

	var quantity eos.Asset
	if val, err := eos.NewAsset(strQuantity); err != nil {
		responseData.State = models.ErrorAssetState
		responseData.Error = err.Error()
		return
	} else {
		quantity = val
	}

	if out, err := models.RootTransfer(to, quantity, memo); err != nil {
		responseData.State = models.ErrorEOSAPIState
		responseData.Error = err.Error()
		return
	} else {
		responseData.State = models.SuccessState
		responseData.Result = out
		return
	}
}

func (c *RestController) ReturnResponsData(remp *models.ResponseData) {
	c.Data["json"] = remp
	c.ServeJSON()
}
