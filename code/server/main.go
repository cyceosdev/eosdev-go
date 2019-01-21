// Package classification EOSIO发行自己的代币
//
// 关于如何使用EOSIO的接口发行自己的代币
//
//      Host: localhost:8000
//      Version: 0.0.1
//
// swagger:meta
package main

import (
	"eosdev-go/code/config"
	"eosdev-go/code/view"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", view.Hello)

	log.Println("启动服务...")
	http.ListenAndServe(config.DefaultApiServer, nil)
}
