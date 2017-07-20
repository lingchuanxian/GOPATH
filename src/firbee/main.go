package main

import (
	_ "firbee/models"
	_ "firbee/routers"
	"github.com/astaxie/beego"
)

const VERSION = "1.0.0"

func main() {
	beego.AppConfig.Set("version", VERSION)
	if beego.BConfig.RunMode == "dev" {
    	beego.BConfig.WebConfig.DirectoryIndex = true
    	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

