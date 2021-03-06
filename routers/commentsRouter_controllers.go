package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Miavega/api_mutants_mercado_libre/controllers:MutantController"] = append(beego.GlobalControllerRouter["github.com/Miavega/api_mutants_mercado_libre/controllers:MutantController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Miavega/api_mutants_mercado_libre/controllers:StatsController"] = append(beego.GlobalControllerRouter["github.com/Miavega/api_mutants_mercado_libre/controllers:StatsController"],
		beego.ControllerComments{
			Method:           "GetStats",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
