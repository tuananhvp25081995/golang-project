package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello/controllers:UserController"] = append(beego.GlobalControllerRouter["hello/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:username`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:UserController"] = append(beego.GlobalControllerRouter["hello/controllers:UserController"],
        beego.ControllerComments{
            Method: "ConfirmEmail",
            Router: `/confirmemail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:UserController"] = append(beego.GlobalControllerRouter["hello/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:UserController"] = append(beego.GlobalControllerRouter["hello/controllers:UserController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
