package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/keesely/nginx"
	"log"
	//"strings"
	//"time"
)

type NginxController struct {
	beego.Controller
}

func ngx() *nginx.Nginx {
	conf, err := config.NewConfig("ini", "/etc/nginx-go.conf")

	if err != nil {
		log.Fatal(err)
	}

	ngx := new(nginx.Nginx)
	ngx.Nginx = conf.String("nginx")
	ngx.Pid = conf.String("pidfile")
	//ngx.Conf_path = conf.String("confpath")
	//ngx.Vhost_path = conf.String("vhostpath")
	return ngx
}

func (this *NginxController) Status() {
	str := "Nginx 进程未运行"
	code := 0

	proc, err := ngx().Status()

	if err == nil {
		code = 200
		str = "Nginx运行中"
	}

	json := map[string]interface{}{
		"code":   code,
		"msg":    str,
		"result": proc,
	}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *NginxController) Test() {
	code := 200
	str := "Nginx 测试成功"

	if _, err := ngx().Test(); err != nil {
		code = 0
		str = "Nginx 测试失败: " + err.Error()
	}

	json := map[string]interface{}{
		"code": code,
		"msg":  str,
	}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *NginxController) Stop() {
	proc, serr := ngx().Status()

	str := ""
	json := map[string]interface{}{
		"code": 0,
		"msg":  "",
	}

	// 进程已经停止
	if proc == nil {
		json["code"] = 200
		json["msg"] = "Nginx已经停止工作"
		this.Data["json"] = json
		this.ServeJSON()
		return
	}

	stop, err := ngx().Stop()

	if false == stop || err != nil {
		if serr != nil {
			str = "Nginx 停止失败: " + err.Error()
			json["code"] = 0
		} else {
			str = fmt.Sprintf("Nginx 停止失败 (%v)", proc.PID)
			json["code"] = 0
		}
	} else {
		str = "Nginx停止工作"
		json["code"] = 200
	}

	json["msg"] = str
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *NginxController) Start() {
	json := map[string]interface{}{
		"code":   0,
		"msg":    "unkown",
		"result": nil,
	}

	proc, srr := ngx().Status()
	if proc == nil || srr != nil {

		start, err := ngx().Start()
		proc, _ = ngx().Status()

		if err != nil {
			json["msg"] = err.Error()
		} else if start != true {
			json["msg"] = "Nginx 启动失败"
		} else {
			json["msg"] = "Nginx 启动成功"
			json["code"] = 200
			json["result"] = proc
		}
	} else {
		json["msg"] = "Nginx 已经启动"
		json["code"] = 200
		json["result"] = proc
	}

	this.Data["json"] = json
	this.ServeJSON()
}

func (this *NginxController) Reload() {
	json := map[string]interface{}{
		"code": 0,
		"msg":  "unkown",
	}

	proc, srr := ngx().Status()
	if proc != nil && srr == nil {
		reload, err := ngx().Reload()

		if err != nil {
			json["msg"] = err.Error()
		} else if reload != true {
			json["msg"] = "Nginx 重载失败"
		} else {
			json["msg"] = "Nginx 重载成功"
			json["code"] = 200
		}
	} else if srr != nil {
		json["msg"] = srr.Error()
	} else {
		json["msg"] = "Nginx 进程不存在，请使用 /start 启用Nginx进程"
	}

	this.Data["json"] = json
	this.ServeJSON()
}
