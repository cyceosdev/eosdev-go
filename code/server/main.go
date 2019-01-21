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
