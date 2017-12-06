package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	_ "nginx-go/routers"
	"os"
)

func main() {
	conf := flag.String("c", "/etc/nginx-go.conf", "set configuration file (default: etc/nginx-go.conf)")
	flag.Parse()

	if _, err := os.Stat(*conf); os.IsNotExist(err) {
		log.Fatalln(err)
	}

	fmt.Println("Loading Configure file: " + string(*conf))
	beego.LoadAppConfig("ini", *conf)
	beego.Run()
}
