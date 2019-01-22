package models

import (
	"github.com/astaxie/beego"
	"github.com/eoscanada/eos-go"
	"sync"
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
