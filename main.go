package main

import (
	"github.com/astaxie/beego"
	"sysX/initx"
	_ "sysX/routers"
)

func main() {
	initx.InitSysX()
	beego.Run()
}