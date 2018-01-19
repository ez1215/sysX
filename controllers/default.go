package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	log.Println("登录页面")
	c.TplName = "login.tpl"
}
