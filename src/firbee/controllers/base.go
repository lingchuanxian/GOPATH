package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) Rsp(status bool, msg string) {
	b.Data["json"] = &map[string]interface{}{"status": status, "msg": msg}
	b.ServeJSON()
}

func (b *BaseController) RspData(status bool, msg string,data interface{}) {
	b.Data["json"] = &map[string]interface{}{"status": status, "msg": msg,"data":data}
	b.ServeJSON()
}