
 //@APIVersion 1.0.0
 //@Title 测试 API 接口
 //@Description mobile has every tool to get any job done, so codename for the new mobile APIs.
 //@Contact astaxie@gmail.com
package routers

import (
	"firbee/controllers"
	"github.com/astaxie/beego"
)

func init() {
   /* beego.Router("/GetUsers", &controllers.UserController{},"get:GetUsers")
    beego.Router("/AddUser", &controllers.UserController{},"post:AddUser")*/

    ns :=
        beego.NewNamespace("/webservice",
            beego.NSNamespace("/user",
                beego.NSInclude(
                    &controllers.UserController{},
                ),
            ),
        )
    beego.AddNamespace(ns)

}
