package main

import (
	"com.zhaoyin/etoken-go/models"
	_ "com.zhaoyin/etoken-go/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.InitEosSdk()
	beego.Run()
}
