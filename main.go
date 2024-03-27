package main

import (
	"demo_beego/config"
	"demo_beego/controllers"
	_ "demo_beego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	config.Init()

	beego.BConfig.EnableErrorsRender = false
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	beego.Run()
}
