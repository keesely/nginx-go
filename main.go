package main

import (
	_ "nginx-go/routers"
	"github.com/astaxie/beego"
)

func main() {
  //beego.AppConfigPath = "/etc/app.conf"
  //beego.ParseConfig()
  beego.LoadAppConfig("ini", "/etc/nginx-go.conf");
	beego.Run()
}

