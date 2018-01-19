package routers

import (
	"sysX/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{},"post:Login")
	beego.Router("/user/register", &controllers.UserController{},"post:Register")
	beego.Router("/index", &controllers.UserController{},"get:Index")
	beego.Router("/user/welcom", &controllers.UserController{},"get:Welcom")
	beego.Router("/user/loginout", &controllers.UserController{},"get:LoginOut")
	beego.Router("/user/sysinfo", &controllers.UserController{},"post:SysInfo")
	beego.Router("/user/list", &controllers.UserController{},"get:List")
	beego.Router("/user/addpage", &controllers.UserController{},"get:AddPage")
	beego.Router("/user/deleteuser", &controllers.UserController{},"post:DeleteUser")

}
