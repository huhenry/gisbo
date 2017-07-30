package routers

import (
	"gisbo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/login", &controllers.IndexController{}, "POST:Login")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signin", &controllers.LoginController{}, "POST:Signin")
	beego.Router("/login", &controllers.LoginController{}, "Get:Get")
}
