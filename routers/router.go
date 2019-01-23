package routers

import (
	"com.zhaoyin/eosdev-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/OpenAccount", &controllers.RestController{}, "post:OpenAccount")
	beego.Router("/api/TestApi", &controllers.RestController{}, "post:TestApi")
	beego.Router("/api/create_token", &controllers.RestController{}, "put:CreateToken")
	beego.Router("/api/get_account", &controllers.RestController{}, "get:GetAccount")
}
