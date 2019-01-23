package models

import (
	"sync"

	"github.com/astaxie/beego"
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/token"
)

var (
	eosurl    string
	pk        string
	account   eos.AccountName
	eosapi    *eos.API
	requerMap sync.Map
)

func InitEosSdk() error {
	eosurl = beego.AppConfig.String("eos_url")
	pk = beego.AppConfig.String("eos_pk")
	account = eos.AN(beego.AppConfig.String("eos_account"))
	beego.Debug("eos_pk======>", pk)
	eosapi = eos.New(eosurl)
	keybag := eos.NewKeyBag()
	err := keybag.ImportPrivateKey(pk)
	if err != nil {
		return err
	}
	eosapi.SetSigner(keybag)
	keys, _ := keybag.AvailableKeys()
	requerMap.Store(account, keys)
	return nil
}

func GetInfo() (infoResp *eos.InfoResp, err error) {
	infoResp, err = eosapi.GetInfo()
	return
}

func GetEosBalance() (out []eos.Asset, err error) {
	out, err = eosapi.GetCurrencyBalance(account, "EOS", "eosio.token")
	return
}

// account 上必须部署了eosio.token的合约
func NewCreate(issuer eos.AccountName, maxSupply eos.Asset) *eos.Action {
	return &eos.Action{
		Account: account,
		Name:    eos.ActN("create"),
		Authorization: []eos.PermissionLevel{
			{Actor: account, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(token.Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

func CreateToken(maxSupply eos.Asset) (out *eos.PushTransactionFullResp, err error) {
	var issuer = account
	var act = NewCreate(issuer, maxSupply)
	out, err = eosapi.SignPushActions(act)
	return
}

func NewIssue(to eos.AccountName, quantity eos.Asset, memo string) *eos.Action {
	return &eos.Action{
		Account: account,
		Name:    eos.ActN("issue"),
		Authorization: []eos.PermissionLevel{
			{Actor: account, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(token.Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

func IssueToken(to eos.AccountName, quantity eos.Asset, memo string) (out *eos.PushTransactionFullResp, err error) {
	var act = NewIssue(to, quantity, memo)
	out, err = eosapi.SignPushActions(act)
	return
}

func GetAccount(name eos.AccountName) (out *eos.AccountResp, err error) {
	out, err = eosapi.GetAccount(name)
	return
}
