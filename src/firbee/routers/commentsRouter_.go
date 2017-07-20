package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["firbee/controllers:UserController"] = append(beego.GlobalControllerRouter["firbee/controllers:UserController"],
		beego.ControllerComments{
			Method: "AddUser",
			Router: `/AddUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["firbee/controllers:UserController"] = append(beego.GlobalControllerRouter["firbee/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUsers",
			Router: `/GetUsers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["firbee/controllers:UserController"] = append(beego.GlobalControllerRouter["firbee/controllers:UserController"],
		beego.ControllerComments{
			Method: "DelUser",
			Router: `/GetUsers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
