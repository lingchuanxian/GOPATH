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
			Method: "DelUser",
			Router: `/DelUser`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["firbee/controllers:UserController"] = append(beego.GlobalControllerRouter["firbee/controllers:UserController"],
		beego.ControllerComments{
			Method: "EditUser",
			Router: `/EditUser`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["firbee/controllers:UserController"] = append(beego.GlobalControllerRouter["firbee/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserById",
			Router: `/GetUserById`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["firbee/controllers:UserController"] = append(beego.GlobalControllerRouter["firbee/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUsers",
			Router: `/GetUsers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
