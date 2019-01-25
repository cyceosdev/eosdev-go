package models

import "github.com/eoscanada/eos-go"

type GetBalancePojo struct {
	Account eos.AccountName `json:"name"`
	Symbol  string          `json:"symbol"`
	Code    string          `json:"code"`
}
