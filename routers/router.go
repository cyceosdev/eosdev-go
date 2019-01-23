package routers

import (
	"com.zhaoyin/eosdev-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/OpenAccount", &controllers.RestController{}, "post:OpenAccount")
	beego.Router("/api/TestApi", &controllers.RestController{}, "post:TestApi")
	beego.Router("/api/create_token", &controllers.RestController{}, "put:CreateToken")
	beego.Router("/api/issue_token", &controllers.RestController{}, "post:IssueToken")
	beego.Router("/api/get_account", &controllers.RestController{}, "get:GetAccount")
	beego.Router("/api/get_currency_balance", &controllers.RestController{}, "get:GetCurrencyBalance")
	beego.Router("/api/create_account", &controllers.RestController{}, "put:CreateAccount")
	beego.Router("/api/root_transfer", &controllers.RestController{}, "post:RootTransfer")
}
