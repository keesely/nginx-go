package routers

import (
	"github.com/astaxie/beego"
	"nginx-go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/status", &controllers.NginxController{}, "*:Status")
	beego.Router("/test", &controllers.NginxController{}, "*:Test")
	beego.Router("/stop", &controllers.NginxController{}, "*:Stop")
	beego.Router("/start", &controllers.NginxController{}, "*:Start")
	beego.Router("/reload", &controllers.NginxController{}, "*:Reload")
}
