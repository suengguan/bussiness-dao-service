package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:projectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:JobController"],
		beego.ControllerComments{
			Method: "DeleteById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ModuleController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ModuleController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:PodController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:PodController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:PodController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:PodController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "DeleteById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["dao-service/bussiness-dao-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "DeleteAll",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
