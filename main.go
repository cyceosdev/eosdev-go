package main

import (
	"com.zhaoyin/eosdev-go/models"
	_ "com.zhaoyin/eosdev-go/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.InitEosSdk()
	beego.Run()
}
